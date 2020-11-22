package main

import (
	"fmt"
	"time"
)

func handleEventMessage(topic string, message string)  {
	fmt.Println("Handling event message")
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
