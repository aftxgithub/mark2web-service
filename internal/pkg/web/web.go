package web

import (
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/thealamu/mark2web-service/internal/pkg/db"
	m2wlog "github.com/thealamu/mark2web-service/internal/pkg/log"
	"github.com/thealamu/mark2web-service/internal/pkg/mark2web"
)

func Start() int {
	srv := &m2wserver{
		service(),
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

func service() *mark2web.Service {
	return &mark2web.Service{
		DB: &db.FSDatabase{
			BaseDir: os.TempDir(),
		},
	}
}

// logger returns a suitable logger for use in handlers
func logger() *log.Logger {
	return m2wlog.New(getLogLevelFromEnv())
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

// getLastPath returns the last path item in a URL.
// For example, for the URL https://example.com/12345, it returns 12345.
// For a URL with no path, it returns an empty string.
func getLastPath(URL string) string {
	urlObj, err := url.Parse(URL)
	if err != nil {
		return ""
	}
	path := urlObj.Path
	i := strings.LastIndex(path, "/")
	if i == -1 {
		return ""
	}
	return path[i+1:]
}
