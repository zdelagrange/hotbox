package main

import (
	"encoding/json"
	"time"

	"fmt"
	"log"
	"net/http"

	"github.com/d2r2/go-dht"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var state = map[string]float32{
	"temperature": 0,
	"humidity":    0,
}

var lastUpdate = time.Now()

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

	newTime := time.Now()
	diff := newTime.Sub(lastUpdate).Seconds
	log.Println(diff)

	if newTime.Sub(lastUpdate).Seconds() > 10 {
		lastUpdate = newTime
		state["temperature"] = temperature
		state["humidity"] = humidity
	}

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
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	})

	router.HandleFunc("/", Root).Methods("GET")
	router.HandleFunc("/api/humidity", Humidity).Methods("GET")
	router.HandleFunc("/api/temperature", Temperature).Methods("GET")

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", handler))
}
