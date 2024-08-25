package consumer

import (
	"api.ddd/pkgs/message_bus/utils"
	"context"
	"errors"
	"github.com/IBM/sarama"
)

// SetUpConsumer start listener to pool message from kafka
func (c *Client) SetUpConsumer() {
	c.Logger.Infof("Setting Up Consumer Listener")
	go func() {
		for {
			err := c.Client.Consume(context.Background(), c.GetTopics(), c)
			if err != nil {
				c.Logger.Warnw("Error Consuming Event", err)
			}
		}
	}()
}

// GetTopics return a list's topics registered
func (c *Client) GetTopics() []string {
	var topics []string
	for _, consumer := range c.Consumers {
		topics = append(topics, consumer.Topic)
	}
	return topics
}

// GetHandlerByTopic get a function that will e executed to a specific topic
func (c *Client) GetHandlerByTopic(topic string) (IEventHandler, error) {
	for _, consumer := range c.Consumers {
		if consumer.Topic == topic {
			return consumer.Handler, nil
		}
	}
	return nil, errors.New("consumer not found")
}

// ConsumeClaim will be executed to process each messages from de consumer
func (c *Client) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		handler, err := c.GetHandlerByTopic(message.Topic)
		if err != nil {
			c.Logger.Errorf("Consumer Handler Not Found For Topic: %s", message.Topic)
			return errors.New("consumer was not registered")
		}

		errorChan := make(chan error)
		go ConsumeMessage(
			handler,
			&utils.Event{
				Headers: message.Headers,
				Topic:   message.Topic,
				Message: message.Value,
			},
			errorChan,
		)

		err = <-errorChan
		if err != nil {
			c.Logger.Errorf("Consumer Handler Error: %s", err.Error())
			return err
		}
		session.MarkMessage(message, "")
	}
	return nil
}

func ConsumeMessage(handler IEventHandler, event *utils.Event, errorChan chan error) {
	err := handler.OnMessage(event)
	errorChan <- err
}

// Setup will be executed before consumer start consuming messages
func (c *Client) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup will be executed after consumer have consumed messages
func (c *Client) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}
