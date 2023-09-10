package utils

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

const (
	SPLITVIDS_REL_PATH = "./skim/uploads/splitVids"
	UPLOADS_REL_PATH   = "./skim/uploads"
)

func SplitVideo() error {

	_, err := os.Stat(SPLITVIDS_REL_PATH)
	if os.IsNotExist(err) {
		mkdirCmd := exec.Command("mkdir", "splitVids")
		mkdirCmd.Dir = UPLOADS_REL_PATH
		mkdirCmd.Run()
	}

	fmt.Println("created splitVidDir")

	var uploadedFileName string
	filepath.Walk(UPLOADS_REL_PATH, func(path string, info fs.FileInfo, err error) error {
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

	err = runCmd(removeAudio)
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
		"splitVids/output%03d.mp4",
	)

	err = runCmd(splitCmd)
	if err != nil {
		return err
	}

	return nil
}

func runCmd(cmd *exec.Cmd) error {
	wg := sync.WaitGroup{}

	cmd.Dir = UPLOADS_REL_PATH

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for scanner.Scan() {
			line := scanner.Text()
			// Process the line of stdout here
			fmt.Println(line)
		}
	}()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	errScanner := bufio.NewScanner(stderr)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for errScanner.Scan() {
			line := errScanner.Text()
			// Process the line of stdout here
			fmt.Println(line)
		}
	}()

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}
