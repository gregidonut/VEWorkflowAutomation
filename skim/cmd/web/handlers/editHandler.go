package handlers

import (
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type templateData struct {
	SplitVidFilePaths []string
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/edit" {
		http.NotFound(w, r)
		return
	}

	_, err := os.Stat(paths.UPLOADS_PATH)
	if os.IsNotExist(err) {
		http.Error(w, "uploads directory does not exist", http.StatusInternalServerError)
		return
	}

	_, err = os.Stat(paths.SPLITVIDS_REL_PATH)
	if err == nil {
		goto afterSplitting
	}

	if err = utils.SplitVideo(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

afterSplitting:
	files := []string{
		"./skim/ui/html/base.html",
		"./skim/ui/html/pages/edit.html",
	}

	var splitVidFiles []string
	dirEntry, err := os.ReadDir(paths.SPLITVIDS_REL_PATH)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, f := range dirEntry {
		if f.IsDir() {
			continue
		}
		splitVidFiles = append(splitVidFiles, fmt.Sprintf("/static/uploads/splitVids/%s", filepath.Base(f.Name())))
	}

	data := &templateData{
		SplitVidFilePaths: splitVidFiles,
	}

	fmt.Println("**split files:**")
	for _, file := range splitVidFiles {
		fmt.Printf("\t- %s\n", file)
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
