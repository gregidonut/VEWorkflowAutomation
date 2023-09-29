package utils

import (
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

func StitchVids() error {
	_, err := os.Stat(paths.RAW_COMMIT_VIDS_REL_PATH)
	if os.IsNotExist(err) {
		os.Mkdir(paths.RAW_COMMIT_VIDS_REL_PATH, os.ModeDir|os.ModePerm)
	}

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

	sort.Strings(filePaths)

	stitchvids := exec.Command(
		"ffmpeg",
		"-f",
		"concat",
		"-safe",
		"0",
		"-i",
		filePaths[len(filePaths)-1],
		"-c",
		"copy",
		fmt.Sprintf("rawCommitVids/output%s.mp4", strings.TrimSuffix(filePaths[len(filePaths)-1], "input.txt")),
	)

	err = RunCmd(stitchvids, paths.WORKSPACE_REL_PATH)
	if err != nil {
		return err
	}

	return nil
}
