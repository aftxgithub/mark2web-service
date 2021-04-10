package mark2web

import (
	"bytes"

	"github.com/gomarkdown/markdown"
)

// MarkdownToHTML returns the HTML equivalent of the passed in markdown
func MarkdownToHTML(md []byte) []byte {
	return bytes.TrimSpace(markdown.ToHTML(md, nil, nil))
}
