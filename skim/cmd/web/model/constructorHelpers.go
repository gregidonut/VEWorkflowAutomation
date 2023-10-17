package model

import (
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/model/osvid"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils/paths"
	"golang.org/x/sync/errgroup"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func (m *Model) ProbeForUploadedVidLength() error {
	m.app.Info("probing for upload vid length")
	defer m.app.Info("finished probing for upload vid length")

	ffprobeCmd := exec.Command(
		"ffprobe",
		m.UploadedVidPath,
		"-show_entries",
		"format=duration",
		"-of",
		"default=noprint_wrappers=1:nokey=1",
	)

	lastLine, err := m.RunCmd(ffprobeCmd)
	if err != nil {
		return err
	}

	splitStrs := strings.Split(lastLine, ".")
	seconds, err := strconv.Atoi(splitStrs[0])
	if err != nil {
		return err
	}

	m.UploadedVidLengthInSeconds = seconds
	return nil
}

func (m *Model) GenInitialOSVids() error {
	m.app.Info("generating initial vids")
	defer m.app.Info("finished generating initial vids")

	m.app.Info("checking for presence if split vids directory")
	_, err := os.Stat(paths.SPLITVIDS_REL_PATH)
	if os.IsNotExist(err) {
		if err = os.Mkdir(paths.SPLITVIDS_REL_PATH, os.ModeDir|os.ModePerm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	initialVidsNumber := m.UploadedVidLengthInSeconds
	m.app.Info(fmt.Sprintf("assigning initial number of videos to be split: %d", initialVidsNumber))
	if m.UploadedVidLengthInSeconds > 30 {
		m.app.Info(fmt.Sprintf("changed initial number of videos to be split: %d", 29))
		initialVidsNumber = 29
	}

	var eg errgroup.Group
	for i := 0; i < initialVidsNumber; i++ {
		outputPath := filepath.Join(paths.SPLITVIDS_REL_PATH, fmt.Sprintf("output_%06d.mp4", i))
		timeStamp := fmt.Sprintf("00:00:%02d", i)

		eg.Go(func() error {
			m.app.Info(fmt.Sprintf("spawning go routine for %s, for timestamp %s", outputPath, timeStamp))
			ffmpegCmd := exec.Command(
				"ffmpeg",
				"-ss",
				timeStamp,
				"-i",
				m.UploadedVidPath,
				"-t",
				"1.0",
				"-an",
				"-vf",
				"scale=360:trunc(ow/a/2)*2,setsar=1",
				"-c:v",
				"libx264",
				"-crf",
				"22",
				outputPath,
			)
			_, err = m.RunCmd(ffmpegCmd)
			if err != nil {
				return err
			}

			m.app.Info(fmt.Sprintf("constructing osvid record for %s, for timestamp %s", outputPath, timeStamp))
			osv, err := osvid.NewOSVid(outputPath, timeStamp)
			if err != nil {
				return err
			}
			m.app.Info(fmt.Sprintf("appending osvid record to model field for %s, for timestamp %s", outputPath, timeStamp))
			m.OSVids = append(m.OSVids, osv)

			return nil
		})
	}

	m.app.Info("finished spawning goroutines for initial split vids!")
	if err = eg.Wait(); err != nil {
		return err
	}

	sort.Slice(m.OSVids, func(i, j int) bool {
		return m.OSVids[i].TimeStampFromUploadedVid < m.OSVids[j].TimeStampFromUploadedVid
	})

	return nil
}
