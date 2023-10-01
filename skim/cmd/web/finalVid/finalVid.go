package finalVid

import (
	"bufio"
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/fsvid"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"os"
	"path/filepath"
	"strings"
)

type FinalVid struct {
	VPath        string `json:"vPath"`
	VBasePath    string `json:"vBasePath"`
	LastModified string `json:"lastModified"`
}

func NewFinalVid() (*FinalVid, error) {
	payload := new(FinalVid)

	payload.VPath = paths.FINAL_VID_PATH
	payload.VBasePath = filepath.Base(paths.FINAL_VID_PATH)

	fsVids, err := fsvid.GenerateFsVidList()
	if err != nil {
		return payload, err
	}

	filePath := strings.TrimSuffix(paths.FINAL_VID_PATH, ".mp4") + ".txt"
	// Open the file for writing, create if it doesn't exist, and truncate if it does
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return payload, err
	}
	defer file.Close()

	// Create a bufio.Writer to efficiently write lines to the file
	writer := bufio.NewWriter(file)

	// Iterate over the array of strings and write each line to the file
	for _, fsv := range fsVids {
		pathAsLine := fmt.Sprintf("file 'actualCommitVids/%s'", filepath.Base(fsv.VBasePath)) + "\n"
		_, err := writer.WriteString(pathAsLine)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return payload, err
		}
	}

	// Flush the bufio.Writer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return payload, err
	}

	fmt.Println("Lines have been written to", filePath)

	return payload, nil
}
