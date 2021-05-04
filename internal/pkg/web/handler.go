package web

import (
	"bytes"
	"fmt"
	tpl "html/template"
	"io"
	"net/http"
	"strings"
)

//go:embed static/scaffold.html
var HTMLScaffoldTemplate string

// HTMLScaffoldData defines data for the scaffold template
type HTMLScaffoldData struct {
	Title   string
	Content string
}

// Parse the HTML template for use in handlers
var template = tpl.Must(tpl.New("scaffold.html").Parse(HTMLScaffoldTemplate))

func (s *server) handleRoot(w http.ResponseWriter, r *http.Request) {
	s.logger.Tracef("handling root path request for method %s\n", r.Method)
	if r.Method == http.MethodPost {
		s.handleMarkdownUpload(w, r)
		return
	}
	s.handleURLresolution(w, r)
}

// handleURLresolution resolves a URL, returning the corresponding HTML
func (s *server) handleURLresolution(w http.ResponseWriter, r *http.Request) {
	s.logger.Tracef("handling url resolution for %+v", r)

	id := getLastPath(r.URL.String())
	s.logger.Debugf("url id is '%s'", id)
	// TODO(thealamu): Naive. Validate 'id' is a sha1
	if id == "" {
		s.logger.Error(fmt.Errorf("empty id is invalid"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	HTMLbytes, err := s.service.HTMLFor(id)
	if err != nil {
		s.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = template.Execute(w, &HTMLScaffoldData{id, string(HTMLbytes)})
	if err != nil {
		s.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// handleMarkdownUpload receives a markdown file and returns a URL to it as static HTML
func (s *server) handleMarkdownUpload(w http.ResponseWriter, r *http.Request) {
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
