package handlers

import (
	"html/template"
	"net/http"
	"os"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/edit" {
		http.NotFound(w, r)
		return
	}

	_, err := os.Stat(UPLOADS_PATH)
	if os.IsNotExist(err) {
		http.Error(w, "directory does not exist", http.StatusInternalServerError)
		return
	}

	ts, err := template.ParseFiles(EDIT_PAGE_PATH)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
