package main

import (
	// "encoding/json"

	"fmt"
	"log"
	"net/http"

	"github.com/d2r2/go-dht"
	"github.com/gorilla/mux"
)

// var (
// 	pin           int
// 	stype         string
// 	boostPerfFlag bool
// )

// func init() {
// 	flag.IntVar(&pin, "pin", 4, "pin")
// 	flag.StringVar(&stype, "sensor-type", "dht22", "sensor type (dht22)")
// 	flag.BoolVar(&boostPerfFlag, "boost", false, "boost performance")
// }

// Root does Root
func Root(w http.ResponseWriter, r *http.Request) {
	temperature, humidity, retried, err := dht.ReadDHTxxWithRetry(dht.DHT22, 4, false, 10)
	if err != nil {
		log.Fatal(err)
	}
	// temperature64 := float64(temperature)
	// humidity64 := float64(humidity)
	fmt.Fprintf(w, `temp: %v\nhumidity: %v\nreried: %v`, temperature, humidity, retried)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Root).Methods("GET")

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", router))
}
