package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/thealamu/mark2web-service/internal/pkg/web"
)

func main() {
	log.SetLevel(log.TraceLevel)

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	serverError := make(chan error, 1)
	shutdownCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go func() {
		serverError <- web.Start()
	}()

	select {
	case err := <-serverError:
		return errors.Wrap(err, "server error")

	case <-shutdownCtx.Done():
		if err := web.Stop(); err != nil {
			return errors.Wrap(err, "could not stop server")
		}
	}

	return nil
}
