package register_services

import (
	"api.ddd/pkgs/message_bus"
	"api.ddd/src/api/server"
	"github.com/IBM/sarama"
)

func ProvideKafkaConnection(logger *server.Logger, config *server.Configuration) message_bus.IMessageBus {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll

	messageBusConfig := message_bus.Configuration{
		Hosts:         []string{config.Kafka.Broker},
		ConsumerGroup: config.Kafka.Group,
		SaramaConfig:  saramaConfig,
	}

	conn, err := message_bus.NewMessageBus(&messageBusConfig)
	if err != nil {
		logger.Errorw("Error Creating Kafka Client", err)
		panic(err)
	}

	return conn
}
