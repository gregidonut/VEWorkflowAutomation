package utilsOld

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

func RunCmd(cmd *exec.Cmd, cmdDir string) error {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current Working Directory: %s\n", currentWorkingDirectory)

	wg := sync.WaitGroup{}

	cmd.Dir = cmdDir

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
