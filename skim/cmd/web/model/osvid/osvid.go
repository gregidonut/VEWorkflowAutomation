package osvid

import "path/filepath"

type OSVid struct {
	VPath                    string
	VBasePath                string
	TimeStampFromUploadedVid string
}

func NewOSVid(vpath, timestamp string) (*OSVid, error) {
	payload := new(OSVid)
	payload.VPath = vpath
	payload.VBasePath = filepath.Base(vpath)
	payload.TimeStampFromUploadedVid = timestamp

	return payload, nil
}
