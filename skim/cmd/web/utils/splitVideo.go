package utils

import (
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func SplitVideo() error {

	_, err := os.Stat(paths.SPLITVIDS_REL_PATH)
	if os.IsNotExist(err) {
		mkdirCmd := exec.Command("mkdir", "splitVids")
		mkdirCmd.Dir = paths.UPLOADS_PATH
		mkdirCmd.Run()
	}

	fmt.Println("created splitVidDir")

	var uploadedFileName string
	filepath.Walk(paths.UPLOADS_PATH, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.Contains(path, "splitVids") {
			return nil
		}

		uploadedFileName = filepath.Base(path)
		return nil
	})

	fmt.Printf("uploaded filename: %q\n", uploadedFileName)

	removeAudio := exec.Command(
		"ffmpeg",
		"-i",
		uploadedFileName,
		"-c:v",
		"copy",
		"-an",
		"-crf",
		"22",
		fmt.Sprintf("%s_no_sound.mp4", strings.TrimSuffix(uploadedFileName, filepath.Ext(uploadedFileName))),
	)

	err = runCmd(removeAudio, paths.UPLOADS_PATH)
	if err != nil {
		return err
	}

	splitCmd := exec.Command(
		"ffmpeg",
		"-i",
		fmt.Sprintf("%s_no_sound.mp4", strings.TrimSuffix(uploadedFileName, filepath.Ext(uploadedFileName))),
		"-c:v",
		"libx264",
		"-crf",
		"22",
		"-map",
		"0",
		"-segment_time",
		"1",
		"-reset_timestamps",
		"1",
		"-g",
		"30",
		"-sc_threshold",
		"0",
		"-force_key_frames",
		`expr:gte(t,n_forced*1)`,
		"-f",
		"segment",
		fmt.Sprintf(
			"splitVids/%s",
			strings.TrimSuffix(uploadedFileName, filepath.Ext(uploadedFileName)),
		)+"_part_%04d.mp4",
	)

	err = runCmd(splitCmd, paths.UPLOADS_PATH)
	if err != nil {
		return err
	}

	return nil
}
