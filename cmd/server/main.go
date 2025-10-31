package main

import (
	"github.com/chetanjangir0/oneresumeplease/internal/api"
	"log"
)

func main() {

	// TODO: use env variables
	cfg := api.Config{
		Addr: ":8080",
	}

	app := &api.Application{
		Config: cfg,
	}

	router := app.Mount()

	log.Fatal(app.Run(router))
}
