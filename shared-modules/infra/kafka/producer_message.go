package kafka

import (
	"errors"

	"github.com/IBM/sarama"
	"github.com/ngoctb13/seta-train/shared-modules/utils"
)

type ProducerMessageOption struct {
	Key       string
	Partition *int32
	Headers   map[string]string
}

var generateUIDFunc = utils.Generate

func prepareProducerMessage(topic string, payload []byte, opt ProducerMessageOption) (*sarama.ProducerMessage, error) {
	if topic == "" {
		return nil, errors.New("topic can't be empty")
	}

	if opt.Key == "" {
		opt.Key = generateUIDFunc()
	}

	pm := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(payload),
		Key:   sarama.StringEncoder(opt.Key),
	}

	if l := len(opt.Headers); l > 0 {
		pm.Headers = make([]sarama.RecordHeader, 0, l)
		for k, v := range opt.Headers {
			pm.Headers = append(pm.Headers, sarama.RecordHeader{
				Key:   []byte(k),
				Value: []byte(v),
			})
		}
	}

	return pm, nil
}
