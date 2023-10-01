package finalVid

import (
	"bufio"
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/fsvid"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func stitchFSVids() error {
	stitchvids := exec.Command(
		"ffmpeg",
		"-f",
		"concat",
		"-safe",
		"0",
		"-i",
		strings.TrimSuffix(filepath.Base(paths.FINAL_VID_PATH), ".mp4")+".txt",
		"-c",
		"copy",
		strings.TrimSuffix(filepath.Base(paths.FINAL_VID_PATH), ".mp4")+"_new.mp4",
	)

	err := utils.RunCmd(stitchvids, filepath.Dir(paths.FINAL_VID_PATH))
	if err != nil {
		return err
	}

	return nil
}

func generateInputFile() error {
	fsVids, err := fsvid.GenerateFsVidList()
	if err != nil {
		return err
	}

	filePath := strings.TrimSuffix(paths.FINAL_VID_PATH, ".mp4") + ".txt"
	// Open the file for writing, create if it doesn't exist, and truncate if it does
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	// Create a bufio.Writer to efficiently write lines to the file
	writer := bufio.NewWriter(file)

	// Iterate over the array of strings and write each line to the file
	for _, fsv := range fsVids {
		// called from finalVid dir path
		pathAsLine := fmt.Sprintf("file '../actualCommitVids/%s'", filepath.Base(fsv.VBasePath)) + "\n"
		_, err := writer.WriteString(pathAsLine)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return err
		}
	}

	// Flush the bufio.Writer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return err
	}

	fmt.Println("Lines have been written to", filePath)
	return nil
}
