package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	Id        int  `json:"id"`
}

var users = []User{}

func (a *app) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// encode users slice to json
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
func (a *app) createUserHandler(w http.ResponseWriter, r *http.Request) {

	// decode request body
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u := User{
		FirstName: payload.FirstName,
		Id:        payload.Id,
	}
	users = append(users, u)
	w.WriteHeader(http.StatusCreated)
}

type app struct {
	addr string
}

func (a *app) postUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("created user"))
}

func main() {
	app := &app{addr: ":8080"}

	// initialze the servermux
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    app.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", app.getUserHandler)
	mux.HandleFunc("POST /users", app.postUserHandler)
	srv.ListenAndServe()
}
