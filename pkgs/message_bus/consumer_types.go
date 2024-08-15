package message_bus

import (
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type ClientConsumer struct {
	Consumers []*Consumer
	Client    sarama.ConsumerGroup
	Logger    *zap.SugaredLogger
}

type Consumer struct {
	Topic   string
	Handler IEventHandler
}

// IEventHandler interface that execute when a message arrive
type IEventHandler interface {
	OnMessage(event *Event)
}

type IConsumer interface {
	SetUpConsumer()
}
