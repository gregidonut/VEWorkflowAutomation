package applicationOld

import (
	"encoding/json"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/model/fsvid"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils/paths"
	"net/http"
	"os"
)

func (app *Application) GetFSVid(w http.ResponseWriter, r *http.Request) {
	_, err := os.Stat(paths.FSVIDS_REL_PATH)
	if os.IsNotExist(err) {
		json.NewEncoder(w).Encode(nil)
		return
	}

	fsVids, err := fsvid.GenerateFsVidList()
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "applicationOld/json")

	json.NewEncoder(w).Encode(fsVids)
}