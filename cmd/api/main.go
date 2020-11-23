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

// Reading GET latest reading
func Reading(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./hotbox.db")
	checkErr(err)

	rows, err := db.Query("SELECT temperature, humidity FROM reading ORDER BY rowid DESC LIMIT 1;")
	checkErr(err)

	var temperature, humidity float32
	if rows.Next() {
		err = rows.Scan(&temperature, &humidity)
		checkErr(err)
	}
	var reading = map[string]float32{
		"temperature": temperature,
		"humidity":    humidity,
	}
	payload, err := json.Marshal(reading)
	checkErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	})

	router.HandleFunc("/api/reading", Reading).Methods("GET")

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", handler))
}
