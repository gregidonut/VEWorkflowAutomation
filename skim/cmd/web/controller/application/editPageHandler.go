package application

import (
	"html/template"
	"net/http"
)

type templateData struct {
	Video string
}

func (app *Application) EditPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/edit" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./skim/ui/html/base.gohtml",
		"./skim/ui/html/pages/edit.gohtml",
		"./skim/ui/html/partials/editPageComponents/initialTimeLineSection.gohtml",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Clear-Site-Data", `"cache"`)

	err = ts.ExecuteTemplate(w, "base", app.model.OSVids)
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}
}
