package message_bus

import (
	"api.ddd/pkgs/mediator"
	"api.ddd/pkgs/message_bus/consumer"
	"api.ddd/pkgs/message_bus/producer"
	"api.ddd/pkgs/message_bus/utils"
	"github.com/IBM/sarama"
)

func NewMessageBus(config *Configuration) (IMessageBus, error) {
	logger := utils.NewLogger()
	producerClient, err := sarama.NewAsyncProducer(config.Hosts, config.SaramaConfig)
	if err != nil {
		logger.Warnw("Failed To Start Sarama Producer", err)
		return nil, err
	}

	consumerClient, err := sarama.NewConsumerGroup(config.Hosts, config.ConsumerGroup, config.SaramaConfig)
	if err != nil {
		logger.Warnw("Failed To Start Sarama Consumer", err)
		return nil, err
	}

	messageBus := &MessageBus{
		consumerClient: consumerClient,
		producerClient: producerClient,
		logger:         logger,
		consumers:      make([]*consumer.Consumer, 0),
		producers:      make([]*producer.Producer, 0),
	}

	return messageBus, nil
}

func (m *MessageBus) WitProducer(topic string, eventType interface{}) IMessageBus {
	newProducer := &producer.Producer{
		Topic:     topic,
		EventType: mediator.TypeOf(eventType),
	}
	m.producers = append(m.producers, newProducer)
	return m
}

// WithConsumer register a consumer to pool messages
func (m *MessageBus) WithConsumer(topic string, handler consumer.IEventHandler) IMessageBus {
	newConsumer := &consumer.Consumer{
		Topic:   topic,
		Handler: handler,
	}
	m.consumers = append(m.consumers, newConsumer)
	return m
}

// Build create final user instance
func (m *MessageBus) Build() IMessageBusClient {
	consumerClient := &consumer.Client{
		Client:    m.consumerClient,
		Logger:    m.logger,
		Consumers: m.consumers,
	}

	producerClient := &producer.Client{
		Client:    m.producerClient,
		Logger:    m.logger,
		Producers: m.producers,
	}

	// SetUp Consumers
	consumerClient.SetUpConsumer()

	return &MessageBusClient{
		consumerClient: consumerClient,
		producerClient: producerClient,
	}
}

func (m MessageBusClient) SendMessage(eventType *utils.Event) {
	m.producerClient.SendMessage(eventType)
}
