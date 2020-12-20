package mqtt

import (
	"fmt"
	paho "github.com/eclipse/paho.paho.golang"
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
	mqttClient paho.Client
}

func (c mqttConnection) Subscribe(topic string, callback func(string, string)) {
	token := c.mqttClient.Subscribe(topic, 0, func(client paho.Client, msg paho.Message) {
		callback(msg.Topic(), string(msg.Payload()))
	})
	if token.Error() != nil {
		log.Fatal("error subscribing to topic", token.Error())
	}
}

func (c mqttConnection) Publish(topic string, message string) {
	token := c.mqttClient.Publish(topic, 0, false, message)
	if token.Error() != nil {
		log.Fatal("error publishing message", token.Error())
	}
}

func newClientOptions(clientId string, uri *url.URL) *paho.ClientOptions {
	opts := paho.NewClientOptions()
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
		log.Fatalf("Cannot parse paho string url: %s", uri)
		os.Exit(1)
	}
	newUUID, _ := uuid.NewUUID()

	client := paho.NewClient(newClientOptions(newUUID.String(), mqttUrl))
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}

	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return mqttConnection{mqttClient: client}
}
