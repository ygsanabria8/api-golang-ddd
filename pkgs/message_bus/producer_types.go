package message_bus

import (
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type ClientProducer struct {
	Producers []*Producer
	Client    sarama.AsyncProducer
	Logger    *zap.SugaredLogger
}

type Producer struct {
	EventType string
	Topic     string
}

type IProducer interface {
	SendMessage(eventType *Event)
}
