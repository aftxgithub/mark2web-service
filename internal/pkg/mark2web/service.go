package mark2web

import (
	"bytes"

	"github.com/gomarkdown/markdown"
)

// Service implements core logic for converting markdown to URL
type Service struct{}

// MarkdownToURL returns a URL for the given markdown.
// It orchestrates converting markdown to HTML,
// generating the URL using the host and storing a mapping from the URL to the HTML.
func (m *Service) MarkdownToURL(md []byte, host string) string {
	return ""
}

// markdownToHTML returns the HTML equivalent of the passed in markdown
func markdownToHTML(md []byte) []byte {
	return bytes.TrimSpace(markdown.ToHTML(md, nil, nil))
}
