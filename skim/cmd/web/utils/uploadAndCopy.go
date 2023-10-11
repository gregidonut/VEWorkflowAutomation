package utils

import (
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils/appInterface"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/utils/paths"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func UploadAndCopy(sourceMPF multipart.File, header *multipart.FileHeader, app appInterface.AppInterface) error {
	app.Debug("Ensure the directory exists, create it if necessary...")
	if _, err := os.Stat(paths.UPLOADS_PATH); os.IsNotExist(err) {
		if err = os.Mkdir(paths.UPLOADS_PATH, os.ModeDir|os.ModePerm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// Create a new file with the same name as the original filename in the upload directory
	newEmptyFile, err := os.Create(filepath.Join(paths.UPLOADS_PATH, header.Filename))
	if err != nil {
		return err
	}
	defer newEmptyFile.Close()

	if err = copyWithProgress(sourceMPF, header.Size, newEmptyFile, app); err != nil {
		return err
	}
	return nil
}

type progressReader struct {
	R      io.Reader
	OnRead func(n int)
}

func (pr *progressReader) Read(p []byte) (n int, err error) {
	n, err = pr.R.Read(p)
	if pr.OnRead != nil {
		pr.OnRead(n)
	}
	return n, err
}

func copyWithProgress(sourceMPF multipart.File, sourceFileSize int64, targetFile *os.File, app appInterface.AppInterface) error {
	app.Debug("starting copyWithProgress function... ")
	// Initialize a progress tracker
	var progressMutex sync.Mutex
	var bytesRead int64
	go func() {
		for {

			progressMutex.Lock()
			app.Info(fmt.Sprintf("Copying... %d bytes", bytesRead))

			appFileProgress := app.CpUploadFileProgressPercentage(int64(float64(bytesRead) / float64(sourceFileSize) * 100))
			app.Info(fmt.Sprintf("copy progress: %d%%", appFileProgress))

			progressMutex.Unlock()

			app.Debug(fmt.Sprintf("total file size of uploaded file: %d", sourceFileSize))
			if bytesRead >= sourceFileSize {
				return
			}
			// Sleep or update interval as per your needs
			// This will give you updates on progress
			// at a regular interval
			time.Sleep(time.Second)
		}
	}()

	app.Debug("Copying...")

	// Create a TeeReader to track the progress
	readerWithProgress := &progressReader{
		R: sourceMPF,
		OnRead: func(n int) {
			progressMutex.Lock()
			bytesRead += int64(n)
			progressMutex.Unlock()
		},
	}

	_, err := io.Copy(targetFile, readerWithProgress)
	if err != nil {
		return err
	}

	app.Debug("ended copyWithProgress function!")
	return nil
}
