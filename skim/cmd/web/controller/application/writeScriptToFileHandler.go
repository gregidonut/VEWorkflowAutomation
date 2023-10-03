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
		app.catchHandlerErr(w, nil, http.StatusMethodNotAllowed)
		return
	}

	var vidScript VidPathAndScript
	if err := json.NewDecoder(r.Body).Decode(&vidScript); err != nil {
		app.catchHandlerErr(w, err, http.StatusBadRequest)
		return
	}

	file, err := os.OpenFile(vidScript.scriptFilePath(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Write the text to the file
	if _, err = file.WriteString(vidScript.Script); err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
