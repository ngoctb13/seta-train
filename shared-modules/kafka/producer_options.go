package kafka

import "github.com/IBM/sarama"

type AckMode int

const (
	// AckModeNone means don't wait for response
	AckModeNone = AckMode(sarama.NoResponse)

	// AckModeLocal means wait for local commit only
	AckModeLocal = AckMode(sarama.WaitForLocal)

	// AckModeInSync means wait for commit to all in-sync replicas
	AckModeInSync = AckMode(sarama.WaitForAll)
)

type producerConfig struct {
	*sarama.Config
}

type ProducerOption func(*producerConfig)

func ProducerWithAckMode(mode AckMode) ProducerOption {
	return func(c *producerConfig) {
		c.Producer.RequiredAcks = sarama.RequiredAcks(mode)
	}
}

func ProducerWithAutoCreateTopics() ProducerOption {
	return func(c *producerConfig) {
		c.Metadata.AllowAutoTopicCreation = true
	}
}
