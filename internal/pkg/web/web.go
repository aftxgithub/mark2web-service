package web

import (
	"net"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	m2wlog "github.com/thealamu/mark2web-service/internal/pkg/log"
)

func Start() int {
	srv := &m2wserver{
		logger(),
		httpServer(),
	}
	srv.setupRoutes()
	srv.logger.Infof("Serving on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		srv.logger.Errorf("could not start server: %v", err)
		return 1
	}
	return 0
}

// logger returns a suitable logger for use in handlers
func logger() *log.Logger {
	l := m2wlog.New(getLogLevelFromEnv())
	l.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	return l
}

// httpServer returns a simple, configured http server
func httpServer() *http.Server {
	return &http.Server{
		Addr:         getRunAddr(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
}

// getRunAddr returns the address to start the server on.
// If no port in environment, it defaults to 8080.
func getRunAddr() string {
	port := getPortFromEnv()
	if port == "" {
		port = "8080"
	}
	return net.JoinHostPort("", port)
}
