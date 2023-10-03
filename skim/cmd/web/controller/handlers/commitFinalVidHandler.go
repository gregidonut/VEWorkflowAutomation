package handlers

import (
	"encoding/json"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/model/finalVid"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils"
	"net/http"
)

func CommitFinalVid(w http.ResponseWriter, r *http.Request) {
	err := utils.FinalStep()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fv, err := finalVid.NewFinalVid()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fv)
}
