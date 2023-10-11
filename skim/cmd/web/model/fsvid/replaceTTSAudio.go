package fsvid

import (
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utilsOld"
	"os"
	"os/exec"
	"strings"
)

func (fsv *FSVid) ReplaceTTSAudio() error {
	if err := utilsOld.GenerateTTS(fsv.VBasePath); err != nil {
		return err
	}

	fmt.Println("starting remove audio command")
	tempFileName := fmt.Sprintf("%snew.mp4", strings.TrimSuffix(fsv.VPath, ".mp4"))
	removeAudio := exec.Command(
		"ffmpeg",
		"-i",
		fsv.VPath,
		"-c:v",
		"copy",
		"-an",
		"-crf",
		"22",
		tempFileName,
	)
	if err := utilsOld.RunCmd(removeAudio, "."); err != nil {
		return err
	}

	if err := os.Remove(fsv.VPath); err != nil {
		return err
	}

	if err := os.Rename(tempFileName, fsv.VPath); err != nil {
		return err
	}
	fmt.Println("finished removing audio")

	return nil
}
