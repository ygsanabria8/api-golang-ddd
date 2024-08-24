package producer

import (
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type Client struct {
	Producers []*Producer
	Client    sarama.AsyncProducer
	Logger    *zap.SugaredLogger
	AppName   string
}

type Producer struct {
	EventType string
	Topic     string
}

type IProducer interface {
	SendMessage(eventType interface{})
}
