package handlers

import (
	"encoding/json"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"net/http"
	"os"
	"path/filepath"
)

func GetFSVid(w http.ResponseWriter, r *http.Request) {
	_, err := os.Stat(paths.FSVIDS_REL_PATH)
	if os.IsNotExist(err) {
		json.NewEncoder(w).Encode(nil)
		return
	}

	var fileNames []string
	files, err := os.ReadDir(paths.FSVIDS_REL_PATH)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		fileNames = append(fileNames, filepath.Base(f.Name()))
	}
	json.NewEncoder(w).Encode(fileNames)
	w.WriteHeader(http.StatusOK)
}
