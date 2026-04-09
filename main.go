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
	// add basic handler at a root
	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("."))))
	// add readiness endpoint handler

	mux.HandleFunc("/healthz", handleReadinessEndpoint)
	// start the server
	serv.ListenAndServe()
}

// handler function for readiness endpoint
func handleReadinessEndpoint(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(200)
	rw.Write([]byte("OK"))
}
