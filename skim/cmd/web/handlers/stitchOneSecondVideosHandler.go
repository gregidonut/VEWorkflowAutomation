package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)

func StitchOneSecondVideos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var vidPathsToStitch []string
	err := json.NewDecoder(r.Body).Decode(&vidPathsToStitch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = os.Stat(paths.WORKSPACE_REL_PATH)
	if os.IsNotExist(err) {
		os.Mkdir(paths.WORKSPACE_REL_PATH, os.ModeDir|os.ModePerm)
	}

	if err = generateInputFile(vidPathsToStitch); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func generateInputFile(vidPathsToStitch []string) error {
	files, err := os.ReadDir(paths.WORKSPACE_REL_PATH)
	if err != nil {
		return err
	}

	if len(files) < 1 {
		f, err := os.Create(fmt.Sprintf("%s/0000input.txt", paths.WORKSPACE_REL_PATH))
		if err != nil {
			return err
		}
		defer f.Close()

		for _, path := range vidPathsToStitch {
			pathAsLine := fmt.Sprintf("file '%s'", path) + "\n"
			_, err := f.WriteString(pathAsLine)
			if err != nil {
				log.Fatal(err)
			}
		}

	} else {
		var fileNames []string
		for _, file := range files {
			fileNames = append(fileNames, file.Name())
		}

		sort.Strings(fileNames)
		lastFile := fileNames[len(fileNames)-1]

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
			pathAsLine := fmt.Sprintf("file '%s'", path) + "\n"
			_, err := f.WriteString(pathAsLine)
			if err != nil {
				return err
			}
		}
	}

	fmt.Printf("files in %s\n", paths.WORKSPACE_REL_PATH)
	for _, file := range files {
		fmt.Printf("\t%s\n", file.Name())
	}

	return nil
}
