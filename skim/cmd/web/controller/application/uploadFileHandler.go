package application

import (
	"errors"
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils"
	"net/http"
)

func (app *Application) UploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data
	app.Logger.Debug("started running ParseMultipartForm function...")
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		app.catchHandlerErr(w, errors.New(fmt.Sprintf("ParseMultipartForm err: %v", err)), http.StatusBadRequest)
		return
	}

	app.Logger.Debug("started running FormFile function...")
	file, header, err := r.FormFile("filename")
	if err != nil {
		app.catchHandlerErr(w, errors.New(fmt.Sprintf("FormFile err: %v", err)), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if err = utils.UploadAndCopy(file, header, app); err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/static", http.StatusSeeOther)
}
