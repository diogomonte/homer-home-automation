//+build wireinject

package main

import (
	"github.com/google/wire"
	mqtt "github.com/diogomonte/home-automation/mqtt"
)

func InitializeMqttClient(uri string) mqtt.MqttClient {
	wire.Build(mqtt.NewMqttClient)
	return nil
}