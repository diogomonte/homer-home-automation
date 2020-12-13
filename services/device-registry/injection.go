package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)



func InitializeDeviceRegistry() DeviceController {
	var appConfig = make(map[string]string)
	appConfig["MYSQL_USER"] = "root"
	appConfig["MYSQL_ROOT_PASSWORD"] = "root"
	appConfig["MYSQL_ROOT_HOST"] = "mysql"
	appConfig["MYSQL_PROTOCOL"] = "tcp"
	appConfig["MYSQL_PORT"] = "3306"
	appConfig["MYSQL_DATABASE"] = "device_registry"
	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		appConfig["MYSQL_USER"],
		appConfig["MYSQL_ROOT_PASSWORD"],
		appConfig["MYSQL_PROTOCOL"],
		appConfig["MYSQL_ROOT_HOST"],
		appConfig["MYSQL_PORT"],
		appConfig["MYSQL_DATABASE"],
	)
 	db, err := gorm.Open("mysql", mysqlCredentials)
 	if err != nil {
 		log.Fatal("failed to connect to mysql database", err.Error())
	}
	repository := NewDeviceRepository(db)
	controller := NewDeviceController(repository)
	return controller
}
