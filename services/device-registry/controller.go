package main

import "github.com/google/uuid"

type CreateRequest struct {

}

type DeviceController interface {
	Create(request CreateRequest)
	Delete(deviceId string)
	List() []Device
	Find(deviceId string) Device
}

type controller struct {
	deviceRepository DeviceRepository
}

var deviceRepository DeviceRepository

func NewDeviceController(repository DeviceRepository) DeviceController {
	deviceRepository = repository
	return controller{deviceRepository: deviceRepository}
}

func (c controller) Create(request CreateRequest) {
	deviceRepository.Save(Device{Id: uuid.New().String()})
}

func (c controller) List() []Device {
	return deviceRepository.List()
}

func (c controller) Find(deviceId string) Device {
	return deviceRepository.Find(deviceId)
}

func (c controller) Delete(deviceId string) {
	deviceRepository.Delete(deviceId)
}
