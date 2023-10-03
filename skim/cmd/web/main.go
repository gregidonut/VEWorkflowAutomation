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

	fileServer := http.FileServer(http.Dir(paths.STATIC_REL_PATH))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", application.Index)
	mux.HandleFunc("/upload", application.UploadFile)
	mux.HandleFunc("/edit", application.Edit)
	mux.HandleFunc("/stitchOneSecondVideos", application.StitchOneSecondVideos)
	mux.HandleFunc("/listCommittedFiles", application.ListCommittedFiles)
	mux.HandleFunc("/writeScriptToFile", application.WriteScriptToFile)
	mux.HandleFunc("/generateFSVids", application.GenerateFSVid)
	mux.HandleFunc("/getFSVids", application.GetFSVid)
	mux.HandleFunc("/editFSVidScript", application.EditFSVidScript)
	mux.HandleFunc("/deleteFSVid", application.DeleteFSVid)
	mux.HandleFunc("/commitFinalVid", application.CommitFinalVid)

	log.Printf("Starting server on %s\n", DEFAULT_PORT)

	err := http.ListenAndServe(DEFAULT_PORT, mux)
	log.Fatal(err)
}
