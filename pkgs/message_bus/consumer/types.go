package consumer

import (
	"api.ddd/pkgs/message_bus/utils"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type Client struct {
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
	OnMessage(event utils.IEvent) error
}

type IConsumer interface {
	SetUpConsumer()
}
