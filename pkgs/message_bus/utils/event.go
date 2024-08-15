package utils

import "encoding/json"

// Event basic structure that contains the topic to send event and event's data
type Event struct {
	Id      string      `json:"Id"`
	Message interface{} `json:"Message"`
}

func (e *Event) GetMessage(obj interface{}) error {
	bytes, _ := json.Marshal(e.Message)
	err := json.Unmarshal(bytes, &obj)
	if err != nil {
		return err
	}
	return nil
}
