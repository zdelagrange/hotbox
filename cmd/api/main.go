package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

type reading struct {
	Temperature float32
	Humidity    float32
}

// Reading GET latest reading; POST new reading
func Reading(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./hotbox.db")
	checkErr(err)
	var newReading reading

	if r.Method == "GET" {
		rows, err := db.Query("SELECT temperature, humidity FROM reading ORDER BY rowid DESC LIMIT 1;")
		checkErr(err)
		defer rows.Close()

		var temperature, humidity float32
		if rows.Next() {
			err = rows.Scan(&temperature, &humidity)
			checkErr(err)
		}
		newReading.Temperature = temperature
		newReading.Humidity = humidity
		payload, err := json.Marshal(newReading)
		checkErr(err)
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&newReading)
		checkErr(err)
		stmt, err := db.Prepare("INSERT INTO reading(datetime, temperature, humidity) values(datetime('now'),?,?)")
		checkErr(err)
		_, err = stmt.Exec(newReading.Temperature, newReading.Humidity)
		checkErr(err)
	}
}

// Readings GET range of readings
func Readings(w http.ResponseWriter, r *http.Request) {
	// var idays int
	// var ihours int
	// var iminutes int
	// var iseconds int
	now := time.Now()

	// query := r.URL.Query()
	// days, ok := query["days"]
	// if !ok || len(days) < 1 {
	// 	http.Error(w, "malformed url", http.StatusBadRequest)
	// } else if len(days) > 0 {
	// 	idays = convertQueryString(days)
	// } else {
	// 	idays = 1
	// }
	// now.AddDate(0, 0, idays)
	// hours, ok := query["hours"]
	// if !ok {
	// 	http.Error(w, "malformed url", http.StatusBadRequest)
	// } else if len(hours) > 0 {
	// 	ihours = convertQueryString(hours)
	// } else {
	// 	ihours = 0
	// }
	// now.Add(time.Duration(ihours))
	// minutes, ok := query["minutes"]
	// if !ok {
	// 	http.Error(w, "malformed url", http.StatusBadRequest)
	// } else if len(minutes) > 0 {
	// 	iminutes = convertQueryString(minutes)
	// } else {
	// 	iminutes = 0
	// }
	// now.Add(time.Duration(iminutes))
	// seconds, ok := query["seconds"]
	// if !ok || len(seconds) < 1 {
	// 	http.Error(w, "malformed url", http.StatusBadRequest)
	// } else if len(seconds) > 0 {
	// 	iseconds = convertQueryString(seconds)
	// } else {
	// 	iseconds = 0
	// }
	// now.Add(time.Duration(iseconds))

	now.AddDate(0, 0, -1)

	var (
		humidity    string
		temperature string
	)
	var readings []*reading
	db, err := sql.Open("sqlite3", "./hotbox.db")
	checkErr(err)

	rows, err := db.Query(fmt.Sprintf("SELECT humidity, temperature FROM reading WHERE datetime > '%s';", now.String()))
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&humidity, &temperature)
		newReading := new(reading)
		fhumidity, _ := strconv.ParseFloat(humidity, 32)
		ftemperature, _ := strconv.ParseFloat(temperature, 32)
		newReading.Humidity = float32(fhumidity)
		newReading.Temperature = float32(ftemperature)
		readings = append(readings, newReading)
	}
	payload, _ := json.Marshal(readings)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func convertQueryString(original []string) int {
	negativeSign := "-"
	negativeSign = string(negativeSign)
	new := strings.Join(original, "")
	new = negativeSign + new
	newInt, _ := strconv.Atoi(new)
	return newInt
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	router.HandleFunc("/api/reading", Reading).Methods("GET", "POST")
	router.HandleFunc("/api/readings", Readings).Methods("GET")

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", handler))
}
