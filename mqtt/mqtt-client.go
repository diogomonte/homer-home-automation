package mqtt

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

func (c mqttConnection) Subscribe(topic string, callback func(string, string))  {
	token := c.mqttClient.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		callback(msg.Topic(), string(msg.Payload()))
	})
	if token.Error() != nil {
		log.Fatal("error subscribing to topic", token.Error())
	}
}

func (c mqttConnection) Publish(topic string, message string)  {
	token := c.mqttClient.Publish(topic, 0, false, message)
	if token.Error() != nil {
		log.Fatal("error publishing message", token.Error())
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

func Connect(uri string) MqttClient {
	mqttUrl, err := url.Parse(uri)
	if err != nil {
		log.Fatalf("Cannot parse mqtt string url: %s", uri)
		os.Exit(1)
	}
	newUUID, _ := uuid.NewUUID()

	client := mqtt.NewClient(newClientOptions(newUUID.String(), mqttUrl))
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}

	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return mqttConnection{mqttClient: client}
}