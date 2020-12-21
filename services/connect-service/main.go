package main

import (
	"encoding/json"
	"fmt"
	"github.com/diogomonte/home-automation/services/shared/mqtt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var mqttClient mqtt.Connection

func handleEventMessage(topic string, message string) {
	fmt.Println("handling event message")

	m, err := mqtt.ParseMqttMessage(message)
	if err != nil {
		fmt.Errorf("error parsing mqtt message %s", message)
	}
	fmt.Print(m)
}

func handleActionRequest(response http.ResponseWriter, r *http.Request) {
	var mqttMessageFormat mqtt.Message

	err := json.NewDecoder(r.Body).Decode(&mqttMessageFormat)
	if err != nil {
		log.Println("error parsing http request body")
		response.WriteHeader(http.StatusBadRequest)
	}

	responseBody, err := json.Marshal(mqttMessageFormat)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
	}

	params := mux.Vars(r)
	deviceId := params["deviceId"]

	if mqttClient != nil {
		mqttClient.Publish("homeautomation/"+deviceId+"/action", string(responseBody))
		response.WriteHeader(200)
	}
}

func main() {
	log.Println("-- Running Connect Service --")

	mqttClient = mqtt.Connect("tcp://mqtt_broker:1883")
	mqttClient.Subscribe("homeautomation/+/event", handleEventMessage)

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/device/{deviceId}/action", handleActionRequest).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", r))
}
