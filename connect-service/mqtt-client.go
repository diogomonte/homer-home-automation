package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"log"
	"net/url"
	"os"
	"time"
)

type MqttClient interface {
	Subscribe(topic string, callback func(string, string))
	Publish(topic string, message string)
}

type mqttConnection struct {
	mqttClient mqtt.Client
}

func NewMqttClient(uri string) MqttClient  {
	mqttUrl, err := url.Parse(uri)
	if err != nil {
		log.Fatalf("Cannot parse mqtt string url: %s", uri)
		os.Exit(1)
	}
	newUUID, _ := uuid.NewUUID()
	client := connect(newUUID.String(), mqttUrl)
	return mqttConnection{mqttClient: client}
}

func (c mqttConnection) Subscribe(topic string, callback func(string, string))  {
	token := c.mqttClient.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
		callback(msg.Topic(), string(msg.Payload()))
	})
	if token.Error() != nil {
		_ = fmt.Errorf("error subscribing to mqtt topic %s", topic)
	}
}

func (c mqttConnection) Publish(topic string, message string)  {
	log.Printf("publishing message on %s", topic)
	token := c.mqttClient.Publish(topic, 0, false, message)
	if token.Error() != nil {
		_ = fmt.Errorf("error publishing mqtt message")
	}
}


func newClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func connect(clientId string, uri *url.URL) mqtt.Client {
	client := mqtt.NewClient(newClientOptions(clientId, uri))
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}