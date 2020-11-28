package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"time"
)

type MqttMessageHeader struct {
	MessageId string	`json:"message_id"`
}

type MqttMessage struct {
	Header MqttMessageHeader 	`json:"header"`
	Body map[string]string		`json:"body"`
}

func main() {
	mqttClient := NewMqttClient("tcp://localhost:1883")
	for true {
		random := strconv.Itoa(rand.Intn((25 - 20) + 20))
		body := make(map[string]string)
		body["temperature"] = random

		message := MqttMessage{
			Header: MqttMessageHeader{
				MessageId: uuid.New().String(),
			},
			Body: body,
		}
		json, _ := json.Marshal(message)
		mqttClient.Publish("homeautomation/soft-device/event", string(json))

		time.Sleep(time.Second * 5)
	}
}
