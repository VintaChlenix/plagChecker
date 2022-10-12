package app

import (
	"io"
	"net/http"
	"os"
	"plagChecker/pkg/checker"
)

func (a *App) UploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

	case "POST":
		a.uploadFile(w, r)
	}
}

func (a *App) uploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		a.log.Errorf("failed to retrieve file: %w", err)
		return
	}
	defer file.Close()

	// Create file
	dst, err := os.Create(handler.Filename)
	defer dst.Close()
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	a.log.Infoln("File Uploaded Successfully")
	norm, err := checker.Normalize(dst)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	a.log.Infoln(norm)

	dst.Seek(0, 0)
	sum, err := checker.GetSum(dst)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	a.log.Infoln(sum)

}
