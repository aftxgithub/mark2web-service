package web

import (
	"io"
	"net/http"
	"os"
	"strings"
)

// handleMarkdownUpload receives a markdown file and returns a URL to it as static HTML
func (s *m2wserver) handleMarkdownUpload(w http.ResponseWriter, r *http.Request) {
	s.logger.Tracef("handling markdown upload for %+v", r)

	if r.Method != http.MethodPost {
		s.logger.Debug("method not allowed. This is a POST endpoint")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ct := r.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "multipart/form-data") {
		s.logger.Debug("required content-type is multipart/form-data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	r.ParseMultipartForm(1 << 10)
	file, _, err := r.FormFile("file")
	if err != nil {
		s.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
