package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = utils.StitchVids(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func generateInputFile(vidPathsToStitch []string) error {
	var files []string
	filepath.Walk(paths.WORKSPACE_REL_PATH, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if strings.Contains(path, "rawCommitVids") {
			return nil
		}

		files = append(files, filepath.Base(path))

		return nil
	})

	if len(files) < 1 && len(vidPathsToStitch) > 1 {
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
		sort.Strings(files)
		lastFile := files[len(files)-1]

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

	fmt.Printf("files in %s\n", paths.WORKSPACE_REL_PATH)
	for _, file := range files {
		fmt.Printf("\t%s\n", file)
	}

	return nil
}
