package common

import (
	"encoding/json"
)

type MqttMessageHeader struct {
	MessageId string	`json:"message_id"`
}

type MqttMessage struct {
	Header MqttMessageHeader 	`json:"header"`
	Body map[string]string		`json:"body"`
}


func ParseMqttMessage(m string) (MqttMessage, error) {
	var mqttMessage MqttMessage
	err := json.Unmarshal([]byte(m), &mqttMessage)
	if err != nil {
		return mqttMessage, err
	}
	return mqttMessage, nil
}