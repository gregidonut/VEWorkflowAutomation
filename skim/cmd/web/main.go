package main

import (
	"log"
	"net/http"
)

const (
	DEFAULT_PORT = ":8080"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	log.Printf("Starting server on %d\n", DEFAULT_PORT)
	err := http.ListenAndServe(DEFAULT_PORT, mux)
	log.Fatal(err)
}
