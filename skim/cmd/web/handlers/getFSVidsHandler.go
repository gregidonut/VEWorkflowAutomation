package handlers

import (
	"encoding/json"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils"
	"net/http"
	"os"
)

func GetFSVid(w http.ResponseWriter, r *http.Request) {
	_, err := os.Stat(paths.FSVIDS_REL_PATH)
	if os.IsNotExist(err) {
		json.NewEncoder(w).Encode(nil)
		return
	}

	fsVids, err := utils.GenerateFsVidList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fsVids)
}
