package main

import (
	"net/http"
)

type app struct {
	addr string
}

func (a *app) getUserHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("user list..."))
}

func (a *app) postUserHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("created user"))
}

func main() {
	app := &app{addr: ":8080"}

	// initialze the servermux
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr: app.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", app.getUserHandler)
	mux.HandleFunc("POST /users", app.postUserHandler)
	srv.ListenAndServe()
}
