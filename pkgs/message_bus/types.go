package message_bus

import (
	"api.ddd/pkgs/message_bus/consumer"
	"api.ddd/pkgs/message_bus/producer"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type IMessageBus interface {
	WithAppName(appName string) IMessageBus
	WithConsumer(topic string, handler consumer.IEventHandler) IMessageBus
	WitProducer(topic string, eventType interface{}) IMessageBus
	Build() IMessageBusClient
}

type MessageBus struct {
	consumerClient sarama.ConsumerGroup
	producerClient sarama.AsyncProducer
	logger         *zap.SugaredLogger
	consumers      []*consumer.Consumer
	producers      []*producer.Producer
	appName        string
}

type MessageBusClient struct {
	consumerClient *consumer.Client
	producerClient *producer.Client
}

type IMessageBusClient interface {
	SendMessage(event interface{})
}

// Configuration struct set required data to connect with a kafka server
type Configuration struct {
	Hosts         []string
	ConsumerGroup string
	SaramaConfig  *sarama.Config
}
