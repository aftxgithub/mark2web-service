package web

import (
	"net"
	"net/http"
	"time"
)

func Start() int {
	srv := &http.Server{
		Addr:         getRunAddr(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		return 1
	}
	return 0
}

func getRunAddr() string {
	return net.JoinHostPort("", getPort())
}
