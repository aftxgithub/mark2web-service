package web

import (
	"net/http"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/thealamu/mark2web-service/internal/pkg/mark2web"
)

type server struct {
	service *mark2web.Service
	logger  *log.Logger
	*http.Server
}

func newServer(addr string, opts ...func(*server) error) (*server, error) {
	s := server{}
	// set sensible server defaults
	s.logger = log.New()
	s.Server = &http.Server{
		Addr: addr,
	}

	for _, opt := range opts {
		err := opt(&s)
		if err != nil {
			return nil, errors.Wrap(err, "could not create server")
		}
	}

	return &s, nil
}

// setupRoutes registers server handlers
// func (s *m2wserver) setupRoutes() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", s.handleRoot)
// 	s.Handler = mux
// }
