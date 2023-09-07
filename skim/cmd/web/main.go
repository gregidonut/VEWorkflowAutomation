package main

import (
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/handlers"
	"log"
	"net/http"
)

const (
	DEFAULT_PORT = ":8080"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/upload", handlers.UploadFile)
	mux.HandleFunc("/edit", handlers.Edit)
	log.Printf("Starting server on %s\n", DEFAULT_PORT)
	err := http.ListenAndServe(DEFAULT_PORT, mux)
	log.Fatal(err)
}
