package web

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (s *m2wserver) handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.handleMarkdownUpload(w, r)
		return
	}
	s.handleURLresolution(w, r)
}

// handleURLresolution resolves a URL, returning the corresponding HTML
func (s *m2wserver) handleURLresolution(w http.ResponseWriter, r *http.Request) {
	s.logger.Tracef("handling url resolution for %+v", r)

	id := getLastPath(r.URL.String())
	s.logger.Debugf("url id is %s", id)
	HTMLbytes, err := s.service.HTMLFor(id)
	if err != nil {
		s.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, string(HTMLbytes))
}

// handleMarkdownUpload receives a markdown file and returns a URL to it as static HTML
func (s *m2wserver) handleMarkdownUpload(w http.ResponseWriter, r *http.Request) {
	s.logger.Tracef("handling markdown upload for %+v", r)

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
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, url)
}
