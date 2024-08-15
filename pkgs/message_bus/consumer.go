package message_bus

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/IBM/sarama"
)

// SetUpConsumer start listener to pool message from kafka
func (c *ClientConsumer) SetUpConsumer() {
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
func (c *ClientConsumer) GetTopics() []string {
	var topics []string
	for _, consumer := range c.Consumers {
		topics = append(topics, consumer.Topic)
	}
	return topics
}

// GetHandlerByTopic get a function that will e executed to a specific topic
func (c *ClientConsumer) GetHandlerByTopic(topic string) (IEventHandler, error) {
	for _, consumer := range c.Consumers {
		if consumer.Topic == topic {
			return consumer.Handler, nil
		}
	}
	return nil, errors.New("consumer not found")
}

// DecodeEvent receive message in bytes and return and object with the message
func (c *ClientConsumer) DecodeEvent(bytes []byte) (*Event, error) {
	event := &Event{}
	err := json.Unmarshal(bytes, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

// ConsumeClaim will be executed to process each messages from de consumer
func (c *ClientConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		handler, err := c.GetHandlerByTopic(message.Topic)
		if err != nil {
			c.Logger.Errorf("Consumer Handler Not Found For Topic: %s", message.Topic)
			return errors.New("consumer was not registered")
		}

		event, err := c.DecodeEvent(message.Value)
		if err != nil {
			return err
		}

		go handler.OnMessage(event)
		session.MarkMessage(message, "")
	}
	return nil
}

// Setup will be executed before consumer start consuming messages
func (c *ClientConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup will be executed after consumer have consumed messages
func (c *ClientConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}
