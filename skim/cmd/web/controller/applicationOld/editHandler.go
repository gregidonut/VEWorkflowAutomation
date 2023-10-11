package applicationOld

import (
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utilsOld"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type templateData struct {
	SplitVidFilePaths []string
}

func (app *Application) Edit(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/edit" {
		http.NotFound(w, r)
		return
	}

	_, err := os.Stat(paths.UPLOADS_PATH)
	if os.IsNotExist(err) {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	_, err = os.Stat(paths.SPLITVIDS_REL_PATH)
	if err == nil {
		goto afterSplitting
	}

	if err = utilsOld.SplitVideo(); err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

afterSplitting:
	files := []string{
		"./skim/ui/html/base.gohtml",
		"./skim/ui/html/pages/edit.gohtml",
	}

	var splitVidFiles []string
	dirEntry, err := os.ReadDir(paths.SPLITVIDS_REL_PATH)
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
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
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}
}
