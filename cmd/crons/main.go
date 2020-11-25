package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/d2r2/go-dht"
	_ "github.com/mattn/go-sqlite3"
)

func UpdateSensorRow() {
	for {
		time.Sleep(60 * time.Second)

		db, err := sql.Open("sqlite3", "./hotbox.db")
		checkErr(err)

		stmt, err := db.Prepare("INSERT INTO reading(datetime, temperature, humidity) values(datetime('now'),?,?)")
		checkErr(err)

		temperature, humidity, _, err := dht.ReadDHTxxWithRetry(dht.DHT22, 4, false, 10)
		_, err = stmt.Exec(temperature, humidity)
		checkErr(err)
		// var reading = map[string]float32{
		// 	"temperature": temperature,
		// 	"humidity":    humidity,
		// }
		// payload, err := json.Marshal(reading)
		// checkErr(err)

		// send request to post api
		// resp, err := http.Post("http://localhost:3000", "application/json", bytes.NewBuffer(payload))
		// checkErr(err)
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
