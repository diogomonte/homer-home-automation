package mqtt

import (
	"encoding/json"
)

type MessageHeader struct {
	MessageId string `json:"message_id"`
}

type Message struct {
	Header MessageHeader     `json:"header"`
	Body   map[string]string `json:"body"`
}

func ParseMqttMessage(m string) (Message, error) {
	var mqttMessage Message
	err := json.Unmarshal([]byte(m), &mqttMessage)
	if err != nil {
		return mqttMessage, err
	}
	return mqttMessage, nil
}
