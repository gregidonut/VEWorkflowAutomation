package application

import (
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"html/template"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Clear-Site-Data", `"cache"`)

	_, err := os.Stat(paths.UPLOADS_PATH)
	if !os.IsNotExist(err) {
		err = os.RemoveAll(paths.UPLOADS_PATH)
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
