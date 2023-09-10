package handlers

import (
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils"
	"html/template"
	"net/http"
	"os"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/edit" {
		http.NotFound(w, r)
		return
	}

	_, err := os.Stat(paths.UPLOADS_PATH)
	if os.IsNotExist(err) {
		http.Error(w, "directory does not exist", http.StatusInternalServerError)
		return
	}

	if err = utils.SplitVideo(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	files := []string{
		"./skim/ui/html/base.html",
		"./skim/ui/html/pages/edit.html",
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
