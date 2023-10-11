package applicationOld

import (
	"encoding/json"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/model/fsvid"
	"net/http"
)

func (app *Application) EditFSVidScript(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.catchHandlerErr(w, nil, http.StatusMethodNotAllowed)
		return
	}

	var fsv fsvid.FSVid
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&fsv); err != nil {
		app.catchHandlerErr(w, err, http.StatusBadRequest)
		return
	}

	if err := fsv.EditScript(); err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	if err := fsv.ReplaceTTSAudio(); err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	if err := fsv.CombineWithTTSAudio(); err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
