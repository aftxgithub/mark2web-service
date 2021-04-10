package main

import (
	"log"

	"github.com/thealamu/mark2web-service/internal/pkg/web"
)

func main() {
	if err := web.Start(); err != nil {
		log.Fatal(err)
	}
}
