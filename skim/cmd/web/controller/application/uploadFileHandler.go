package application

import (
	"errors"
	"fmt"
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func (app *Application) UploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data
	app.Logger.Debug("started running ParseMultipartForm function...")
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		app.catchHandlerErr(w, errors.New(fmt.Sprintf("ParseMultipartForm err: %v", err)), http.StatusBadRequest)
		return
	}

	app.Logger.Debug("started running FormFile function...")
	file, header, err := r.FormFile("filename")
	if err != nil {
		app.catchHandlerErr(w, errors.New(fmt.Sprintf("FormFile err: %v", err)), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	app.Logger.Debug("Ensure the directory exists, create it if necessary...")
	if _, err = os.Stat(paths.UPLOADS_PATH); os.IsNotExist(err) {
		if err = os.Mkdir(paths.UPLOADS_PATH, os.ModeDir|os.ModePerm); err != nil {
			app.catchHandlerErr(w, err, http.StatusInternalServerError)
			return
		}
	} else if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	// Create a new file with the same name as the original filename in the upload directory
	newEmptyFile, err := os.Create(filepath.Join(paths.UPLOADS_PATH, header.Filename))
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}
	defer newEmptyFile.Close()

	if err = copyWithProgress(file, header.Size, newEmptyFile, app); err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/static", http.StatusSeeOther)
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

func copyWithProgress(sourceMPF multipart.File, sourceFileSize int64, targetFile *os.File, app *Application) error {
	app.Logger.Debug("starting copyWithProgress function... ")
	// Initialize a progress tracker
	var progressMutex sync.Mutex
	var bytesRead int64
	go func() {
		for {
			progressMutex.Lock()
			app.Logger.Info(fmt.Sprintf("Copying... %d bytes", bytesRead))

			app.copyUploadFileProgressMutex.Lock()
			app.copyUploadFileProgressPercentage = int64(float64(bytesRead) / float64(sourceFileSize) * 100)
			app.Logger.Info(fmt.Sprintf("copy progress: %d%%", app.copyUploadFileProgressPercentage))
			app.copyUploadFileProgressMutex.Unlock()

			progressMutex.Unlock()

			app.Logger.Debug(fmt.Sprintf("total file size of uploaded file: %d", sourceFileSize))
			if bytesRead >= sourceFileSize {
				return
			}
			// Sleep or update interval as per your needs
			// This will give you updates on progress
			// at a regular interval
			time.Sleep(time.Second)
		}
	}()

	app.Logger.Debug("Copying...")

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

	app.Logger.Debug("ended copyWithProgress function!")
	return nil
}
