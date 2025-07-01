package main

import (
	"github.com/prok05/spot-instrument-service/config"
	"github.com/prok05/spot-instrument-service/internal/app"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
