package mark2web

import (
	"bytes"

	"github.com/gomarkdown/markdown"
)

// M2WService implements core logic for converting markdown to URL
type M2WService struct{}

// URLFor returns a URL for the given markdown.
// It orchestrates converting markdown to HTML,
// generating the URL and storing a mapping from the URL to the HTML.
func (m *M2WService) URLFor(markdown []byte) string {
}

// markdownToHTML returns the HTML equivalent of the passed in markdown
func markdownToHTML(md []byte) []byte {
	return bytes.TrimSpace(markdown.ToHTML(md, nil, nil))
}
