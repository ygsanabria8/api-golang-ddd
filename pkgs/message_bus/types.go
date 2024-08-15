package message_bus

import (
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type IMessageBus interface {
	WithConsumer(topic string, handler IEventHandler) IMessageBus
	WitProducer(topic string, eventType interface{}) IMessageBus
	Build() IMessageBusClient
}

type MessageBus struct {
	consumerClient sarama.ConsumerGroup
	producerClient sarama.AsyncProducer
	logger         *zap.SugaredLogger
	consumers      []*Consumer
	producers      []*Producer
}

type MessageBusClient struct {
	consumerClient *ClientConsumer
	producerClient *ClientProducer
}

type IMessageBusClient interface {
	SendMessage(eventType *Event)
}

// Configuration struct set required data to connect with a kafka server
type Configuration struct {
	Hosts         []string
	ConsumerGroup string
	SaramaConfig  *sarama.Config
}
