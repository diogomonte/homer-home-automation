package main

import (
	"github.com/jinzhu/gorm"
)

type Device struct {
	Id string `gorm:"primary_key"`
}

type DeviceRepository interface {
	Save(device Device)
	Delete(deviceId string)
	List() []Device
	Find(deviceId string) Device
}

type repository struct {
	db *gorm.DB
}

func NewDeviceRepository(deviceDb *gorm.DB) DeviceRepository {
	if deviceDb.HasTable(&Device{}) == false {
		deviceDb.CreateTable(&Device{})
	}
	return repository{db: deviceDb}
}

func (r repository) Save(device Device) {
	r.db.Save(device)
}

func (r repository) List() []Device {
	var devices []Device
	r.db.Find(&devices)
	return devices
}

func (r repository) Delete(deviceId string) {
	var device Device
	r.db.Delete(&device, deviceId)
}

func (r repository) Find(deviceId string) Device {
	var device Device
	r.db.First(&device, deviceId)
	return device
}
