package application

import (
	"fmt"
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

	var data []templateData
	for i := 0; i < 30; i++ {
		data = append(data, templateData{
			Video: fmt.Sprintf("ina_mo_%d.mp4", i),
		})
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}
}
