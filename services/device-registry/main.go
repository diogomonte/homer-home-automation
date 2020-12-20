package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var deviceController DeviceController

func HandleCreateAction(res http.ResponseWriter, req *http.Request) {
	var request CreateRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	deviceController.Create(request)
}

func HandleListDevices(w http.ResponseWriter, _ *http.Request) {
	devices := deviceController.List()
	responseBody, _ := json.Marshal(devices)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(responseBody))
}

func HandleFindDevice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	device := deviceController.Find(params["deviceId"])
	responseBody, _ := json.Marshal(device)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(responseBody))
}

func HandleDeleteDevice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	devices := deviceController.Find(params["deviceId"])
	responseBody, _ := json.Marshal(devices)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(responseBody))
}

func main() {
	log.Print("-- Start Device Registry --")
	deviceController = InitializeDeviceRegistry()
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/devices", HandleCreateAction).Methods("POST")
	r.HandleFunc("/devices/{deviceId}", HandleDeleteDevice).Methods("DELETE")
	r.HandleFunc("/devices/{deviceId}", HandleFindDevice).Methods("GET")
	r.HandleFunc("/devices", HandleListDevices).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", r))
}
