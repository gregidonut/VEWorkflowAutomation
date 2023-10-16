package model

import (
	"os/exec"
	"strconv"
	"strings"
)

func (m *Model) ProbeForUploadedVidLength() error {
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

func (m *Model) GenInitial30OSVids() error {
	return nil
}
