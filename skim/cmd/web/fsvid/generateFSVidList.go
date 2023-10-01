package fsvid

import (
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func GenerateFsVidList() ([]FSVid, error) {
	files, err := os.ReadDir(paths.FSVIDS_REL_PATH)
	if err != nil {
		return nil, nil
	}

	var fsVids []FSVid
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		scriptBasePath := strings.TrimSuffix(filepath.Base(f.Name()), ".mp4") + ".txt"
		fsVid, err := NewFSVid(map[string]string{
			"vPath":      fmt.Sprintf("%s/%s", paths.FSVIDS_REL_PATH, filepath.Base(f.Name())),
			"scriptPath": fmt.Sprintf("%s/%s", paths.RAW_COMMIT_VIDS_REL_PATH, scriptBasePath),
		})
		if err != nil {
			return nil, nil
		}

		fsVids = append(fsVids, *fsVid)
	}

	sort.Slice(fsVids, func(i, j int) bool {
		return fsVids[i].VBasePath < fsVids[j].VBasePath
	})

	return fsVids, nil
}
