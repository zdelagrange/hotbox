package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/d2r2/go-dht"
	_ "github.com/mattn/go-sqlite3"
)

type reading struct {
	Temperature float32
	Humidity    float32
}

func UpdateSensorRow() {
	var newReading reading
	for {
		time.Sleep(60 * time.Second)

		temperature, humidity, _, err := dht.ReadDHTxxWithRetry(dht.DHT22, 4, false, 10)
		checkErr(err)
		newReading.Humidity = humidity
		newReading.Temperature = temperature
		payload, err := json.Marshal(newReading)
		checkErr(err)

		resp, err := http.Post("http://localhost:3000", "application/json", bytes.NewBuffer(payload))
		if resp.StatusCode != 200 {
			payload, _ := ioutil.ReadAll(resp.Request.Body)
			log.Printf("Request Failed: Status Code {%d} | Request URL {%s} | Request Payload {%s}", resp.StatusCode, resp.Request.URL, payload)
		}
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	UpdateSensorRow()
}
