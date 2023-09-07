package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

const (
	SPLITVIDS_REL_PATH = "../../../uploads/splitVids"
)

func SplitVideo() error {
	wg := sync.WaitGroup{}

	_, err := os.Stat(SPLITVIDS_REL_PATH)
	if os.IsNotExist(err) {
		mkdirCmd := exec.Command("mkdir", SPLITVIDS_REL_PATH)
		mkdirCmd.Run()
	}

	//{{
	ffmpegCmd := exec.Command(
		"ffmpeg",
		"-i",
		"krombopulos_michael.mp4",
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
	ffmpegCmd.Dir = "../../../uploads"

	stdout, err := ffmpegCmd.StdoutPipe()
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

	stderr, err := ffmpegCmd.StderrPipe()
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

	err = ffmpegCmd.Start()
	if err != nil {
		return err
	}

	err = ffmpegCmd.Wait()
	if err != nil {
		return err
	}

	wg.Wait()
	//}}

	return nil
}
