package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type VidPathAndScript struct {
	Script  string `json:"script"`
	VidPath string `json:"vidPath"`
}

func (vpas *VidPathAndScript) scriptFilePath() string {
	name := strings.TrimSuffix(vpas.VidPath, filepath.Ext(vpas.VidPath))

	return fmt.Sprintf("skim/ui%s.txt", name)

}

func (app *Application) WriteScriptToFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var vidScript VidPathAndScript
	err := json.NewDecoder(r.Body).Decode(&vidScript)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, err := os.OpenFile(vidScript.scriptFilePath(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Write the text to the file
	_, err = file.WriteString(vidScript.Script)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
