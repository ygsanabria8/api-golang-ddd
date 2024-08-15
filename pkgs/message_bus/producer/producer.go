package producer

import (
	"api.ddd/pkgs/mediator"
	"api.ddd/pkgs/message_bus/utils"
	"encoding/json"
	"errors"
	"github.com/IBM/sarama"
)

func (p *Client) FindTopicByProducer(eventType interface{}) (string, error) {
	eventTypeString := mediator.TypeOf(eventType)
	for _, producer := range p.Producers {
		if producer.EventType == eventTypeString {
			return producer.Topic, nil
		}
	}
	return "", errors.New("consumer not found")
}

func (p *Client) SendMessage(eventType *utils.Event) {
	go func() {
		topic, err := p.FindTopicByProducer(eventType.Message)
		if err != nil {
			p.Logger.Errorw("Error Getting Topic By Event", err)
			return
		}

		messageBytes, err := json.Marshal(eventType)
		if err != nil {
			p.Logger.Errorw("Error Serializing message", err)
			return
		}

		message := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(messageBytes),
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
