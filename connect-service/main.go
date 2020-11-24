package main

import (
	"fmt"
	"time"
)

func handleEventMessage(topic string, message string)  {
	fmt.Println("Handling event message")
	_, err := ParseMqttMessage(message)
	if err != nil {

	}
}

func handleActionMessage(topic string, message string)  {
	fmt.Println("Handling action message")
}

func main()  {
	mqtt := InitializeMqttClient("tcp://localhost:1883")
	mqtt.Subscribe("homeautomation/+/event", handleEventMessage)
	mqtt.Subscribe("homeautomation/+/action", handleActionMessage)

	for true {
		time.Sleep(time.Second)
	}
}
