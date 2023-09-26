package handlers

import (
	"encoding/json"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils"
	"net/http"
	"os"
)

func FSVids(w http.ResponseWriter, r *http.Request) {
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

	//files, err := os.ReadDir(paths.FSVIDS_REL_PATH)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//for _, file := range files {
	//	if file.IsDir() {
	//		continue
	//	}
	//
	//	fileNames = append(fileNames, file.Name())
	//}
	//sort.Strings(fileNames)

	//json.NewEncoder(w).Encode(fileNames)
	w.WriteHeader(http.StatusOK)
}
