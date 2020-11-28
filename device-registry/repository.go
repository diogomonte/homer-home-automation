package main

import (
	"github.com/jinzhu/gorm"
)

type Device struct {
	Id string
}

type DeviceRepository interface {
	Save(device Device)
	List() []Device
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

func (r repository) Save(device Device)  {
	r.db.Save(device)
}

func (r repository) List() []Device  {
	var devices []Device
	r.db.Find(&devices)
	return devices
}
