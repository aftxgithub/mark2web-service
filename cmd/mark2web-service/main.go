package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/thealamu/mark2web-service/internal/pkg/web"
)

func main() {
	log.SetLevel(log.TraceLevel)
	os.Exit(web.Start())
}
