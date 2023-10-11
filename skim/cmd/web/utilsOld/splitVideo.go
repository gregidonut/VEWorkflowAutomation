package utilsOld

import (
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func SplitVideo() error {
	_, err := os.Stat(paths.SPLITVIDS_REL_PATH)
	if os.IsNotExist(err) {
		os.Mkdir(paths.SPLITVIDS_REL_PATH, os.ModeDir|os.ModePerm)
	}

	fmt.Println("created splitVidDir")

	var uploadedFileName string
	files, err := os.ReadDir(paths.UPLOADS_PATH)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		uploadedFileName = filepath.Base(f.Name())
	}

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

	err = RunCmd(removeAudio, paths.UPLOADS_PATH)
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

	err = RunCmd(splitCmd, paths.UPLOADS_PATH)
	if err != nil {
		return err
	}

	return nil
}
