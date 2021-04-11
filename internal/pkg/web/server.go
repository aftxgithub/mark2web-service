package web

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/thealamu/mark2web-service/internal/pkg/mark2web"
)

// m2wserver is a http server, holding our run dependencies.
type m2wserver struct {
	service *mark2web.Service
	logger  *log.Logger
	*http.Server
}

// setupRoutes registers server handlers
func (s *m2wserver) setupRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleMarkdownUpload)
	s.Handler = mux
}
