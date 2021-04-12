package web

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
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
	var buf bytes.Buffer
	file, _, err := r.FormFile("file")
	if err != nil {
		s.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	io.Copy(&buf, file)

	url, err := s.service.MarkdownToURL(buf.Bytes(), r.Host)
	if err != nil {
		s.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.logger.Tracef("url for markdown is %s", url)
	fmt.Fprintln(w, url)
}
