package model

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func (m *Model) RunCmd(cmd *exec.Cmd) (string, error) {
	var lastLine string

	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}
	m.app.Debug(fmt.Sprintf("Current Working Directory: %s\n", currentWorkingDirectory))

	wg := sync.WaitGroup{}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	outScanner := bufio.NewScanner(stdout)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for outScanner.Scan() {
			line := outScanner.Text()
			// Process the line of stdout here
			lastLine = line
			m.app.Debug(line)
		}
	}()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	errScanner := bufio.NewScanner(stderr)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for errScanner.Scan() {
			line := errScanner.Text()
			// Process the line of stderr here
			m.app.Debug(line)
		}
	}()

	m.app.Info(fmt.Sprintf("running command: '%s'", cmd))
	err = cmd.Start()
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	wg.Wait()
	return lastLine, err
}
