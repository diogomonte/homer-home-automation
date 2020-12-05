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
	if mqttClient != nil {
		mqttClient.Publish("homeautomation/" + deviceId + "/action", "hello! I am alive: " + deviceId)
	}
	response.WriteHeader(200)
}


func handleGetRequest(response http.ResponseWriter, r *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("hello"))
}

func main()  {
	log.Println("-- Running Connect Service --")
	mqttClient = InitializeMqttClient("tcp://mqtt_broker:1883")
	mqttClient.Subscribe("homeautomation/+/event", handleEventMessage)

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/device/{deviceId}/action", handleActionRequest).Methods("POST")
	r.HandleFunc("/devices", handleGetRequest).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", r))
}
