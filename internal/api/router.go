package api 

import (
	"log"
	"net/http"
	"time"
)

type Application struct{
	Config Config
}

type Config struct{
	Addr string
}	

func (app *Application) Mount() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /health", app.healthCheckHandler)
	return router
}

func (app *Application) Run(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr: app.Config.Addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("server has started at %s", app.Config.Addr) 

	return srv.ListenAndServe()
}
