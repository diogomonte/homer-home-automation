package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var deviceController DeviceController

func HandleCreateAction(res http.ResponseWriter, req *http.Request)  {
	var request CreateRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	deviceController.CreateDevice(request)
}

func HandleListDevices(w http.ResponseWriter, _ *http.Request)  {
	devices := deviceController.ListDevices()
	responseBody, _ := json.Marshal(devices)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(responseBody))
}

func main()  {
	log.Print("-- Start Device Registry --")
	deviceController = InitializeDeviceRegistry()
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/devices", HandleCreateAction).Methods("POST")
	r.HandleFunc("/devices", HandleListDevices).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", r))
}
