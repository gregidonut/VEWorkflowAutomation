package osvid

import (
	"os"
	"path/filepath"
	"time"
)

type OSVid struct {
	VPath                    string
	VBasePath                string
	TimeStampFromUploadedVid string
	LastModified             string
	Committed                bool
}

func NewOSVid(vpath, timestamp string) (*OSVid, error) {
	payload := new(OSVid)
	payload.VPath = vpath
	payload.VBasePath = filepath.Base(vpath)
	payload.TimeStampFromUploadedVid = timestamp
	f, err := os.Stat(vpath)
	if err != nil {
		return payload, err
	}
	payload.LastModified = f.ModTime().Format(time.RFC3339)

	return payload, nil
}
