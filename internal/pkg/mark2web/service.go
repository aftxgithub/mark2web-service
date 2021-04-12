package mark2web

import (
	"bytes"
	"crypto/sha1"
	"fmt"

	"github.com/gomarkdown/markdown"
	"github.com/thealamu/mark2web-service/internal/pkg/db"
)

// Service implements core logic for converting markdown to URL
type Service struct {
	DB db.DB
}

// MarkdownToURL generates a URL for the markdown,
// creates a mapping of the URL to the markdown and returns the URL
func (*Service) MarkdownToURL(md []byte, host string) string {
	HTMLEquiv := markdownToHTML(md)
	path := shasumOf(HTMLEquiv)
	return fmt.Sprintf("%s/%s", host, path)
}

// shasumOf returns the sha1sum of data
func shasumOf(data []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(data))
}

// markdownToHTML returns the HTML equivalent of the passed in markdown
func markdownToHTML(md []byte) []byte {
	return bytes.TrimSpace(markdown.ToHTML(md, nil, nil))
}
