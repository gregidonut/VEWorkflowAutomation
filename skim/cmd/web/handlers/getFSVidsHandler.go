package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/fsvid"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GetFSVid(w http.ResponseWriter, r *http.Request) {
	_, err := os.Stat(paths.FSVIDS_REL_PATH)
	if os.IsNotExist(err) {
		json.NewEncoder(w).Encode(nil)
		return
	}

	var fsVids []fsvid.FSVid
	files, err := os.ReadDir(paths.FSVIDS_REL_PATH)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		scriptBasePath := strings.TrimSuffix(filepath.Base(f.Name()), ".mp4") + ".txt"
		fsVid, err := fsvid.NewFSVid(map[string]string{
			"vPath":      fmt.Sprintf("%s/%s", paths.FSVIDS_REL_PATH, filepath.Base(f.Name())),
			"scriptPath": fmt.Sprintf("%s/%s", paths.RAW_COMMIT_VIDS_REL_PATH, scriptBasePath),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fsVids = append(fsVids, *fsVid)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fsVids)
}
