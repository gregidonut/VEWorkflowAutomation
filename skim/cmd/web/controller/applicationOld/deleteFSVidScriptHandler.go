package applicationOld

import (
	"encoding/json"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/model/fsvid"
	"net/http"
)

func (app *Application) DeleteFSVid(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		app.catchHandlerErr(w, nil, http.StatusMethodNotAllowed)
		return
	}

	var fsv fsvid.FSVid
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&fsv); err != nil {
		app.catchHandlerErr(w, err, http.StatusBadRequest)
		return
	}

	if err := fsv.Delete(); err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
