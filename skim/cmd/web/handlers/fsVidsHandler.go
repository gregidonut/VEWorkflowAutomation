package handlers

import (
	"encoding/json"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils"
	"net/http"
	"os"
)

func GenFSVid(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := utils.CombineFSVidWithTTSAudio()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var fileNames []string
	_, err = os.Stat(paths.FSVIDS_REL_PATH)
	if os.IsNotExist(err) {
		json.NewEncoder(w).Encode(fileNames)
		return
	}

	w.WriteHeader(http.StatusOK)
}
