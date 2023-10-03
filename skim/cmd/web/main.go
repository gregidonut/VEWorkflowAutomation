package main

import (
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/controller/application"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"log"
	"net/http"
)

const (
	DEFAULT_PORT = ":8080"
)

func main() {
	mux := http.NewServeMux()
	app, err := application.NewApplication()
	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.Dir(paths.STATIC_REL_PATH))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.Index)
	mux.HandleFunc("/upload", app.UploadFile)
	mux.HandleFunc("/edit", app.Edit)
	mux.HandleFunc("/stitchOneSecondVideos", app.StitchOneSecondVideos)
	mux.HandleFunc("/listCommittedFiles", app.ListCommittedFiles)
	mux.HandleFunc("/writeScriptToFile", app.WriteScriptToFile)
	mux.HandleFunc("/generateFSVids", app.GenerateFSVid)
	mux.HandleFunc("/getFSVids", app.GetFSVid)
	mux.HandleFunc("/editFSVidScript", app.EditFSVidScript)
	mux.HandleFunc("/deleteFSVid", app.DeleteFSVid)
	mux.HandleFunc("/commitFinalVid", app.CommitFinalVid)

	log.Printf("Starting server on %s\n", DEFAULT_PORT)

	err = http.ListenAndServe(DEFAULT_PORT, mux)
	log.Fatal(err)
}
