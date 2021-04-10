package main

import (
	"os"

	"github.com/thealamu/mark2web-service/internal/pkg/web"
)

func main() {
	os.Exit(web.Start())
}
