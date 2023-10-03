package application

import (
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils"
	"net/http"
)

func (app *Application) GenerateFSVid(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.catchHandlerErr(w, nil, http.StatusMethodNotAllowed)
		return
	}

	err := utils.CombineFSVidWithTTSAudio()
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
