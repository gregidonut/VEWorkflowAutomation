package handlers

import (
	"encoding/json"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"net/http"
	"os"
	"sort"
	"strings"
)

func ListCommittedFiles(w http.ResponseWriter, r *http.Request) {
	var fileNames []string
	w.Header().Set("Content-Type", "application/json")

	_, err := os.Stat(paths.COMMIT_VIDS_REL_PATH)
	if os.IsNotExist(err) {
		json.NewEncoder(w).Encode(fileNames)
		return
	}

	files, err := os.ReadDir(paths.COMMIT_VIDS_REL_PATH)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.Contains(file.Name(), "txt") {
			continue
		}

		fileNames = append(fileNames, file.Name())
	}
	sort.Strings(fileNames)

	json.NewEncoder(w).Encode(fileNames)
}
