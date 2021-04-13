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
		Addr:    addr,
		Handler: getRoutes(&s),
	}
	// init a service
	service, err := mark2web.NewService()
	if err != nil {
		return nil, errors.Wrap(err, "could not create server")
	}
	s.service = service

	// load config opts
	for _, opt := range opts {
		err := opt(&s)
		if err != nil {
			return nil, errors.Wrap(err, "could not create server")
		}
	}

	return &s, nil
}

// getRoutes registers server handlers
func getRoutes(s *server) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleRoot)
	return mux
}
