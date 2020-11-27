#!/bin/bash
go get github.com/d2r2/go-dht
go get github.com/gorilla/mux
go get github.com/bitly/go-simplejson
go get github.com/rs/cors
go get github.com/mattn/go-sqlite3
go build -o cmd/crons/ cmd/crons/main.go
go build -o cmd/api/ cmd/api/main.go
