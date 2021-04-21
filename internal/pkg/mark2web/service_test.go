package mark2web

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/gomarkdown/markdown/parser"
)

func TestShasumOf(t *testing.T) {
	testData := []byte("Markdown Data")
	expected := "fba53b3a2a39f676ab4d1b4ff09e43c8d11729fa"

	got := shasumOf(testData)
	if got != expected {
		t.Errorf("expected '%s' from shasumOf, got '%s'", expected, got)
	}
}

func TestMarkdownToHTML(t *testing.T) {
	testMarkdown, err := ioutil.ReadFile("./testfixtures/markdown.md")
	if err != nil {
		t.Fatal(err)
	}

	testHTML, err := ioutil.ReadFile("./testfixtures/mdhtml.html")
	if err != nil {
		t.Fatal(err)
	}

	extensions := parser.CommonExtensions
	psr := parser.NewWithExtensions(extensions)

	gotHTML := markdownToHTML(testMarkdown, psr)
	if !bytes.Equal(testHTML, gotHTML) {
		t.Fatalf("Rendered HTML does not match expected")
	}
}
