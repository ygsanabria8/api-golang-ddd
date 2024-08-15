package register_services

import (
	"api.ddd/pkgs/message_bus"
	"api.ddd/src/api/server"
	"github.com/IBM/sarama"
)

func ProvideKafkaConnection(logger *server.Logger) message_bus.IMessageBus {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll

	config := message_bus.Configuration{
		Hosts:         []string{"127.0.0.1:9092"},
		ConsumerGroup: "golang-api-ddd",
		SaramaConfig:  saramaConfig,
	}

	conn, err := message_bus.NewMessageBus(&config)
	if err != nil {
		logger.Errorw("Error Creating Kafka Client", err)
		panic(err)
	}

	return conn
}
