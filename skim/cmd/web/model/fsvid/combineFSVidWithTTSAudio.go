package fsvid

import (
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils"
	"os"
	"os/exec"
	"strings"
)

func (fsv *FSVid) CombineWithTTSAudio() error {
	_, err := os.Stat(paths.FSVIDS_REL_PATH)
	if os.IsNotExist(err) {
		os.Mkdir(paths.FSVIDS_REL_PATH, os.ModeDir|os.ModePerm)
	}

	tempFileName := fmt.Sprintf("%snew.mp4", strings.TrimSuffix(fsv.VPath, ".mp4"))
	CombineFSVidWithTTSCmd := exec.Command(
		"ffmpeg",
		"-i",
		fsv.VPath,
		"-i",
		fmt.Sprintf("%s/%s.mp3", paths.RAW_COMMIT_VIDS_REL_PATH, strings.TrimSuffix(fsv.VBasePath, ".mp4")),
		"-c:v",
		"copy",
		"-c:a",
		"aac",
		tempFileName,
	)

	if err = utils.RunCmd(CombineFSVidWithTTSCmd, "."); err != nil {
		return err
	}

	if err = os.Rename(tempFileName, fsv.VPath); err != nil {
		return err
	}

	return nil
}
