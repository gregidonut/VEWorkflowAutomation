package finalVid

import (
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils/paths"
	"os"
	"path/filepath"
	"strings"
	"time"
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

	if err := generateInputFile(); err != nil {
		return payload, err
	}

	if err := stitchFSVids(); err != nil {
		return payload, err
	}

	// renaming because it is always created by the above function as "path.new"
	if err := os.Rename(
		strings.TrimSuffix(paths.FINAL_VID_PATH, ".mp4")+"_new.mp4", paths.FINAL_VID_PATH,
	); err != nil {
		return payload, err
	}

	fInfo, err := os.Stat(paths.FINAL_VID_PATH)
	if err != nil {
		return payload, err
	}
	payload.LastModified = fInfo.ModTime().Format(time.RFC3339)

	return payload, nil
}
