package applicationOld

import (
	"encoding/json"
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func (app *Application) StitchOneSecondVideos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.catchHandlerErr(w, nil, http.StatusMethodNotAllowed)
		return
	}

	var vidPathsToStitch []string
	err := json.NewDecoder(r.Body).Decode(&vidPathsToStitch)
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusBadRequest)
		return
	}

	if _, err = os.Stat(paths.WORKSPACE_REL_PATH); os.IsNotExist(err) {
		if err = os.Mkdir(paths.WORKSPACE_REL_PATH, os.ModeDir|os.ModePerm); err != nil {
			app.catchHandlerErr(w, err, http.StatusInternalServerError)
			return
		}
	}

	if err = generateInputFile(vidPathsToStitch); err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	if err = utils.StitchVids(); err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func generateInputFile(vidPathsToStitch []string) error {
	var filePaths []string
	files, err := os.ReadDir(paths.WORKSPACE_REL_PATH)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		filePaths = append(filePaths, filepath.Base(f.Name()))
	}

	if len(filePaths) < 1 && len(vidPathsToStitch) > 1 {
		f, err := os.Create(fmt.Sprintf("%s/0000input.txt", paths.WORKSPACE_REL_PATH))
		if err != nil {
			return err
		}
		defer f.Close()

		for _, path := range vidPathsToStitch {
			// called from workspace dir
			pathAsLine := fmt.Sprintf("file '../splitVids/%s'", filepath.Base(path)) + "\n"
			_, err := f.WriteString(pathAsLine)
			if err != nil {
				log.Fatal(err)
			}
		}

	} else {
		sort.Strings(filePaths)
		lastFile := filePaths[len(filePaths)-1]

		fileNumber := lastFile[:4]
		fileNumberAsInt, err := strconv.Atoi(fileNumber)
		if err != nil {
			return err
		}

		fileNumberAsInt++

		f, err := os.Create(fmt.Sprintf("%s/%04dinput.txt", paths.WORKSPACE_REL_PATH, fileNumberAsInt))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		for _, path := range vidPathsToStitch {
			// called from workspace dir
			pathAsLine := fmt.Sprintf("file '../splitVids/%s'", filepath.Base(path)) + "\n"
			_, err := f.WriteString(pathAsLine)
			if err != nil {
				return err
			}
		}
	}

	fmt.Printf("filePaths in %s\n", paths.WORKSPACE_REL_PATH)
	for _, file := range filePaths {
		fmt.Printf("\t%s\n", file)
	}

	return nil
}
