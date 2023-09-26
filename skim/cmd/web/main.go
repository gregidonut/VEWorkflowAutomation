package main

import (
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/handlers"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"log"
	"net/http"
)

const (
	DEFAULT_PORT = ":8080"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(paths.STATIC_REL_PATH))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/upload", handlers.UploadFile)
	mux.HandleFunc("/edit", handlers.Edit)
	mux.HandleFunc("/stitchOneSecondVideos", handlers.StitchOneSecondVideos)
	mux.HandleFunc("/listCommittedFiles", handlers.ListCommittedFiles)
	mux.HandleFunc("/writeScriptToFile", handlers.WriteScriptToFile)
	mux.HandleFunc("/fsVids", handlers.FSVids)

	log.Printf("Starting server on %s\n", DEFAULT_PORT)

	err := http.ListenAndServe(DEFAULT_PORT, mux)
	log.Fatal(err)
}
