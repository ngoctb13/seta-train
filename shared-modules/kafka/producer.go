package kafka

import (
	"context"
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/ngoctb13/seta-train/shared-modules/config"
	pkgerrors "github.com/pkg/errors"
)

type Producer struct {
	client       sarama.Client
	syncProducer sarama.SyncProducer
	clientID     string
}

func NewSyncProducer(ctx context.Context, appCfg config.AppConfig, brokers []string, opts ...ProducerOption) (*Producer, error) {
	log.Printf("Initializing kafka producer")

	baseCfg := initBaseConfig(appCfg)
	cfg := &producerConfig{baseCfg}
	cfg.Producer.Return.Successes = true

	for _, opt := range opts {
		opt(cfg)
	}

	client, err := sarama.NewClient(brokers, cfg.Config)
	if err != nil {
		return nil, pkgerrors.Wrap(err, "client init failed")
	}

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return nil, pkgerrors.Wrap(err, "producer init failed")
	}

	sp := &Producer{
		client:       client,
		syncProducer: producer,
		clientID:     cfg.ClientID,
	}

	return sp, nil
}

func initBaseConfig(appCfg config.AppConfig) *sarama.Config {
	cfg := sarama.NewConfig()

	// Use config from AppConfig if available
	if appCfg.Kafka != nil {
		cfg.RackID = appCfg.Kafka.RackID
		cfg.ClientID = appCfg.Kafka.ClientID

		// Parse Kafka version
		if appCfg.Kafka.Version != "" {
			version, err := sarama.ParseKafkaVersion(appCfg.Kafka.Version)
			if err != nil {
				log.Printf("Warning: invalid kafka version %s, using default", appCfg.Kafka.Version)
			} else {
				cfg.Version = version
			}
		}
	} else {
		// Fallback to default values
		cfg.RackID = "localhost"
		cfg.ClientID = "rest-service"
	}

	cfg.Metadata.AllowAutoTopicCreation = false

	// Add retry and timeout settings
	cfg.Producer.Retry.Max = 3
	cfg.Producer.Retry.Backoff = time.Millisecond * 100
	cfg.Producer.Timeout = 5 * time.Second

	return cfg
}

func (p *Producer) SendMessage(ctx context.Context, topic string, payload []byte, opt ProducerMessageOption) (int32, int64, error) {
	start := time.Now()
	defer func() {
		log.Printf("Kafka message sent to topic %s in %v", topic, time.Since(start))
	}()

	pm, err := prepareProducerMessage(topic, payload, opt)
	if err != nil {
		return 0, 0, err
	}

	var partition int32
	var offset int64

	partition, offset, err = p.syncProducer.SendMessage(pm)
	if err != nil {
		return 0, 0, err
	}

	return partition, offset, nil
}

func (p *Producer) Close() error {
	if err := p.syncProducer.Close(); err != nil {
		return pkgerrors.Wrap(err, "could not stop producer")
	}

	if !p.client.Closed() {
		if err := p.client.Close(); err != nil {
			return pkgerrors.Wrap(err, "could not stop producer client")
		}
	}

	return nil
}
