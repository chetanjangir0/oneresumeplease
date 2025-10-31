package api

import (
	"log"
	"net/http"
	"time"

	"github.com/chetanjangir0/oneresumeplease/internal/handler"
	"github.com/chetanjangir0/oneresumeplease/internal/middleware"
)

type Application struct {
	Config Config
}

type Config struct {
	Addr string
}

func (app *Application) Mount() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /health", handler.HealthCheck)

	// global middewares
	stack := middleware.CreateStack(
		middleware.Logger,
		middleware.Recover,
		middleware.CORS,
	)

	return stack(router)
}

func (app *Application) Run(router http.Handler) error {
	srv := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at %s", app.Config.Addr)

	return srv.ListenAndServe()
}
