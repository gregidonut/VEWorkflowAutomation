package handlers

import (
	"html/template"
	"net/http"
	"os"
)

const (
	UPLOADS_DIR = "./skim/uploads"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	_, err := os.Stat(UPLOADS_DIR)
	if !os.IsNotExist(err) {
		err = os.RemoveAll(UPLOADS_DIR)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	files := []string{
		"./skim/ui/html/base.html",
		"./skim/ui/html/pages/index.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
