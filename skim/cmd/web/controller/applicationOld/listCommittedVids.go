package applicationOld

import (
	"encoding/json"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils/paths"
	"net/http"
	"os"
	"sort"
	"strings"
)

func (app *Application) ListCommittedFiles(w http.ResponseWriter, r *http.Request) {
	var fileNames []string
	w.Header().Set("Content-Type", "applicationOld/json")

	_, err := os.Stat(paths.RAW_COMMIT_VIDS_REL_PATH)
	if os.IsNotExist(err) {
		json.NewEncoder(w).Encode(fileNames)
		return
	}

	files, err := os.ReadDir(paths.RAW_COMMIT_VIDS_REL_PATH)
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.Contains(file.Name(), "txt") {
			continue
		}
		if strings.Contains(file.Name(), "mp3") {
			continue
		}

		fileNames = append(fileNames, file.Name())
	}
	sort.Strings(fileNames)

	json.NewEncoder(w).Encode(fileNames)
}
