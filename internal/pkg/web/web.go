package web

import (
	"net"
	"net/http"
	"time"
)

func Start() int {
	runAddr := net.JoinHostPort("", getPort())
	srv := &http.Server{
		Addr:         runAddr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		return 1
	}
	return 0
}
