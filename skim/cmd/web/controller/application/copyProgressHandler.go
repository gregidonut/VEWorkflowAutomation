package application

import (
	"encoding/json"
	"net/http"
)

func (app *Application) CopyProgress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicationOld/json")
	app.CopyUploadFileProgressMutex.Lock()
	if err := json.NewEncoder(w).Encode(app.CopyUploadFileProgressPercentage); err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
	}
	app.CopyUploadFileProgressMutex.Unlock()
}
