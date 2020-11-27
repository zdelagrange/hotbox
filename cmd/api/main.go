package main

import (
	"database/sql"
	"encoding/json"

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

// Reading GET latest reading
func Reading(w http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	db, err := sql.Open("sqlite3", "./hotbox.db")
	checkErr(err)
	newReading := reading{}

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
		//write new reading
		var NewNewReading reading
		// body, err := ioutil.ReadAll(r.Body)
		// checkErr(err)
		// err = json.Unmarshal(body, %NewNewReading)
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&NewNewReading)
		checkErr(err)
		stmt, err := db.Prepare("INSERT INTO reading(datetime, temperature, humidity) values(datetime('now'),?,?)")
		checkErr(err)
		_, err = stmt.Exec(NewNewReading.Temperature, NewNewReading.Humidity)
		checkErr(err)
	}
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
	// router.HandleFunc("/api/readings", Readings).Methods("GET")

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", handler))
}
