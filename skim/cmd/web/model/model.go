package model

import (
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/model/osvid"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils/appInterface"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils/paths"
	"os"
	"path/filepath"
	"strings"
)

// Model is responsible for wrapping all the model objects so that they
// can be neatly bridged over to the main application object
type Model struct {
	UploadedVidPath            string
	UploadedVidLengthInSeconds int
	app                        appInterface.AppInterface
	OSVids                     []*osvid.OSVid
	OSVidsComplete             bool
}

func NewModel(app appInterface.AppInterface) (*Model, error) {
	app.Debug("creating application model..")
	defer app.Debug("finished creating application model!")

	payload := new(Model)
	payload.app = app

	if _, err := os.Stat(paths.UPLOADS_PATH); err != nil {
		return payload, err
	}

	dirEntries, err := os.ReadDir(paths.UPLOADS_PATH)
	if err != nil {
		return payload, err
	}

	for _, de := range dirEntries {
		if de.IsDir() {
			continue
		}

		if !strings.Contains(de.Name(), ".mp4") {
			continue
		}

		payload.UploadedVidPath = filepath.Join(paths.UPLOADS_PATH, de.Name())
	}

	if err = payload.ProbeForUploadedVidLength(); err != nil {
		return payload, err
	}
	app.Info(fmt.Sprintf("length of uploaded video in seconds: %d", payload.UploadedVidLengthInSeconds))

	if err = payload.GenInitialOSVids(); err != nil {
		return payload, err
	}

	return payload, nil
}
