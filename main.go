package main

import (
	"net/http"
)

func main() {
	// create a multiplexer for a server
	mux := http.NewServeMux()
	// create server that uses created mux
	serv := http.Server{
		Addr:		":8080",
		Handler:	mux,
	}
	// start the server
	serv.ListenAndServe()
}
