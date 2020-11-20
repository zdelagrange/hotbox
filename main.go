package main

import (
	"encoding/json"

	"fmt"
	"log"
	"net/http"

	"github.com/d2r2/go-dht"
	"github.com/gorilla/mux"
)

var state = map[string]float32{
	"temperature": 0,
	"humidity":    0,
}

// Root does Root
func Root(w http.ResponseWriter, r *http.Request) {
	temperature, humidity, retried, err := dht.ReadDHTxxWithRetry(dht.DHT22, 4, false, 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, `temp: %v\nhumidity: %v\nreried: %v`, temperature, humidity, retried)
}

// SensorState set and get state
func SensorState() map[string]float32 {
	temperature, humidity, _, err := dht.ReadDHTxxWithRetry(dht.DHT22, 4, false, 10)
	if err != nil {
		log.Fatal(err)
	}

	state["temperature"] = temperature
	state["humidity"] = humidity
	return (state)
}

// Humidity blah
func Humidity(w http.ResponseWriter, r *http.Request) {
	SensorState()

	humidity := map[string]float32{
		"humidity": state["humidity"],
	}

	payload, err := json.Marshal(humidity)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

// Temperature temp
func Temperature(w http.ResponseWriter, r *http.Request) {
	SensorState()

	temperature := map[string]float32{
		"temperature": state["temperature"],
	}

	payload, err := json.Marshal(temperature)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Root).Methods("GET")
	router.HandleFunc("/api/humidity", Humidity).Methods("GET")
	router.HandleFunc("/api/temperature", Temperature).Methods("GET")

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", router))
}
