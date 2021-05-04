package web

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	tmpl "text/template"
)

func TestE2E(t *testing.T) {
	s, err := newServer(":8080")
	if err != nil {
		t.Fatal(err)
	}

	handler := s.Handler

	// Submit markdown
	testMarkdownBytes := []byte("# Markdown Data")

	mPartReader, contentType, err := createMultipart(testMarkdownBytes)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/", mPartReader)
	req.Header.Set("Content-Type", contentType)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected code '%d', got '%d'", http.StatusCreated, rr.Code)
	}

	gotURL, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Get HTML
	id := getLastPath(string(gotURL))

	expectedHTML := `<h1 id="markdown-data">Markdown Data</h1>`

	var buf bytes.Buffer
	template := tmpl.Must(tmpl.New("scaffold.html").ParseFiles("./static/scaffold.html"))
	err = template.Execute(&buf, HTMLScaffoldData{id, expectedHTML})
	if err != nil {
		t.Fatal(err)
	}

	req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", id), nil)
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected code '%d', got '%d'", http.StatusOK, rr.Code)
	}

	gotHTMLBytes, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	expectedHTMLBytes := buf.Bytes()
	if !bytes.Equal(gotHTMLBytes, expectedHTMLBytes) {
		t.Fatalf("expected HTML content to be '%s', got '%s'", expectedHTMLBytes, gotHTMLBytes)
	}
}

func createMultipart(filedata []byte) (io.Reader, string, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// TODO(thealamu): Look into CreateFormFile. Why do I have to provide a filename?
	fw, err := w.CreateFormFile("file", "file")
	if err != nil {
		return nil, "", err
	}

	fw.Write(filedata)
	w.Close()

	return &b, w.FormDataContentType(), nil
}
