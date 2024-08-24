package producer

import (
	"api.ddd/pkgs/mediator"
	"encoding/json"
	"errors"
	"github.com/IBM/sarama"
	"github.com/google/uuid"
)

// FindTopicByEventType return topic name where the message will be sended
func (p *Client) FindTopicByEventType(eventType interface{}) (string, error) {
	eventTypeString := mediator.TypeOf(eventType)
	for _, producer := range p.Producers {
		if producer.EventType == eventTypeString {
			return producer.Topic, nil
		}
	}
	return "", errors.New("consumer not found")
}

// SendMessage send a message to a specific topic
func (p *Client) SendMessage(event interface{}) {
	go func() {
		topic, err := p.FindTopicByEventType(event)
		if err != nil {
			p.Logger.Errorw("Error Getting Topic By Event", err)
			return
		}

		messageBytes, err := json.Marshal(event)
		if err != nil {
			p.Logger.Errorw("Error Serializing message", err)
			return
		}

		message := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(messageBytes),
			Headers: []sarama.RecordHeader{
				{
					Key:   []byte("Id"),
					Value: []byte(uuid.New().String()),
				},
				{
					Key:   []byte("App_Name"),
					Value: []byte(p.AppName),
				},
			},
		}

		p.Client.Input() <- message
		select {
		case err := <-p.Client.Errors():
			p.Logger.Errorw("Error Sending Message To Topic: "+topic+" With Error:", err)
		case <-p.Client.Successes():
			p.Logger.Infow("Sent Message To Topic: " + topic + " Was Successfully")
		}
	}()
}
