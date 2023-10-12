package application

import (
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils/paths"
	"html/template"
	"net/http"
	"os"
)

func (app *Application) IndexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	app.CopyUploadFileProgressMutex.Lock()
	app.CopyUploadFileProgressPercentage = 0
	app.CopyUploadFileProgressMutex.Unlock()

	_, err := os.Stat(paths.UPLOADS_PATH)
	if !os.IsNotExist(err) {
		err = os.RemoveAll(paths.UPLOADS_PATH)
		if err != nil {
			app.catchHandlerErr(w, err, http.StatusInternalServerError)
			return
		}
	}

	files := []string{
		"./skim/ui/html/base.gohtml",
		"./skim/ui/html/pages/index.gohtml",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Clear-Site-Data", `"cache"`)

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}
}