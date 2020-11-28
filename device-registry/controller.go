package main

import "github.com/google/uuid"

type CreateRequest struct {

}

type DeviceController interface {
	CreateDevice(request CreateRequest)
	ListDevices() []Device
}

type controller struct {
	deviceRepository DeviceRepository
}

var deviceRepository DeviceRepository

func NewDeviceController(repository DeviceRepository) DeviceController {
	deviceRepository = repository
	return controller{deviceRepository: deviceRepository}
}

func (c controller) CreateDevice(request CreateRequest) {
	deviceRepository.Save(Device{Id: uuid.New().String()})
}

func (c controller) ListDevices() []Device {
	return deviceRepository.List()
}
