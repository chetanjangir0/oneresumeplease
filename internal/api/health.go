package api 

import "net/http"

func (app *Application) HealthCheckHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Ok"))
}
