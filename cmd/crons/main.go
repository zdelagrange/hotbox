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
		time.Sleep(10 * time.Second)

		db, err := sql.Open("sqlite3", "./hotbox.db")
		checkErr(err)

		stmt, err := db.Prepare("INSERT INTO reading(datetime, temperature, humidity) values(datetime('now'),?,?)")
		checkErr(err)

		temperature, humidity, _, err := dht.ReadDHTxxWithRetry(dht.DHT22, 4, false, 10)
		_, err = stmt.Exec(temperature, humidity)
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
