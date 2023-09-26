package utils

import (
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"io/fs"
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
	sort.Strings(files)

	stitchvids := exec.Command(
		"ffmpeg",
		"-f",
		"concat",
		"-safe",
		"0",
		"-i",
		files[len(files)-1],
		"-c",
		"copy",
		fmt.Sprintf("rawCommitVids/output%s.mp4", strings.TrimSuffix(files[len(files)-1], "input.txt")),
	)

	err = runCmd(stitchvids, paths.WORKSPACE_REL_PATH)
	if err != nil {
		return err
	}

	return nil
}
