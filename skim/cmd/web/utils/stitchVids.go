package utils

import (
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"os"
	"os/exec"
)

func StitchVids() error {
	_, err := os.Stat(paths.COMMIT_VIDS_REL_PATH)
	if os.IsNotExist(err) {
		os.Mkdir(paths.COMMIT_VIDS_REL_PATH, os.ModeDir|os.ModePerm)
	}

	stitchvids := exec.Command(
		"ffmpeg",
		"-f",
		"concat",
		"-safe",
		"0",
		"-i",
		"0000input.txt",
		"-c",
		"copy",
		"commitVids/0000output.mp4",
	)

	err = runCmd(stitchvids, paths.WORKSPACE_REL_PATH)
	if err != nil {
		return err
	}

	return nil
}
