package main

import (
	"log"
	"github.com/chetanjangir0/oneresumeplease/internal/api"
)

func main(){
	cfg:= api.Config{
		Addr: ":8080",
	}

	app := &api.Application{
		Config: cfg,
	}

	router := app.Mount()

	log.Fatal(app.Run(router))
}
