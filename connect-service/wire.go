//+build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeMqttClient(uri string) MqttClient {
	wire.Build(NewMqttClient)
	return nil
}