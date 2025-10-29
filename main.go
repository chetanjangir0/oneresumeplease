package main

import "net/http"

type server struct {
	addr string
}

func (s *server) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}

func main() {
	s := &server{addr: ":8080"}
	http.ListenAndServe(s.addr, s)
}
