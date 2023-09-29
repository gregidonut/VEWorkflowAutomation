package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/fsvid"
	"net/http"
)

func EditFSVidScript(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var fsv fsvid.FSVid
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&fsv); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	if err := fsv.EditScript(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("replacing tts audio..")
	if err := fsv.ReplaceTTSAudio(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("replacing tts audio..")

	w.WriteHeader(http.StatusOK)
}
