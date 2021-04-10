package web

import (
	"net"
	"net/http"
	"time"
)

func Start() error {
	runAddr := net.JoinHostPort("", getPort())
	srv := &http.Server{
		Addr:         runAddr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	return srv.ListenAndServe()
}
