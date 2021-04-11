package web

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// m2wserver is a http server, holding our run dependencies.
type m2wserver struct {
	log *log.Logger
	*http.Server
}
