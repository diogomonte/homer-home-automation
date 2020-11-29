package main

import (
	"fmt"
	mqtt "github.com/diogomonte/home-automation/mqtt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var mqttClient mqtt.MqttClient

func handleEventMessage(topic string, message string)  {
	fmt.Println("Handling event message")
	m, err := mqtt.ParseMqttMessage(message)
	if err != nil {
		fmt.Errorf("error parsing mqtt message %s", message)
	}
	fmt.Print(m)
}

func handleActionRequest(response http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	deviceId := params["deviceId"]
	if mqttClient == nil {
		log.Println("Null mqtt client")
	} else {
		mqttClient.Publish("homeautomation/" + deviceId + "/action", "hello! I am alive: " + deviceId)
	}
}

func main()  {
	log.Println("-- Running Connect Service --")
	mqttClient = InitializeMqttClient("tcp://localhost:1883")
	mqttClient.Subscribe("homeautomation/+/event", handleEventMessage)

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/device/{deviceId}/action", handleActionRequest).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", r))
}
