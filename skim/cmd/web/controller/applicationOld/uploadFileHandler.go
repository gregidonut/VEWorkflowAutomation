package applicationOld

import (
	"github.com/gregidonut/VEWorkflowAutomation/skim/cmd/web/paths"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func (app *Application) UploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("filename")
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Ensure the directory exists, create it if necessary
	if _, err := os.Stat(paths.UPLOADS_PATH); os.IsNotExist(err) {
		if err = os.Mkdir(paths.UPLOADS_PATH, os.ModeDir|os.ModePerm); err != nil {
			app.catchHandlerErr(w, err, http.StatusInternalServerError)
			return
		}
	} else if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	// Create a new file with the same name as the original filename in the upload directory
	uploadedFile, err := os.Create(filepath.Join(paths.UPLOADS_PATH, handler.Filename))
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	// Copy the uploaded file data to the new file
	_, err = io.Copy(uploadedFile, file)
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/edit", http.StatusSeeOther)
}
