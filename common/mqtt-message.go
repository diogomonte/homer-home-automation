package common

type MqttMessageHeader struct {
	MessageId string	`json:"message_id"`
}

type MqttMessage struct {
	Header MqttMessageHeader 	`json:"header"`
	Body map[string]string		`json:"body"`
}
