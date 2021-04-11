package mark2web

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestMarkdownToHTML(t *testing.T) {
	testMarkdown, err := ioutil.ReadFile("./testfixtures/markdown.md")
	if err != nil {
		t.Fatal(err)
	}

	testHTML, err := ioutil.ReadFile("./testfixtures/mdhtml.html")
	if err != nil {
		t.Fatal(err)
	}

	gotHTML := markdownToHTML(testMarkdown)
	if !bytes.Equal(testHTML, gotHTML) {
		t.Fatalf("Rendered HTML does not match expected")
	}
}
