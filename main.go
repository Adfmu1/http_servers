package main

import _ "github.com/lib/pq"

import (
	"net/http"
)


func main() {
	// create a multiplexer for a server
	mux := http.NewServeMux()
	// create server that uses created mux
	serv := &http.Server{
		Addr:		":8080",
		Handler:	mux,
	}
	apiConf := &apiConfig{}
	// add basic handler at a root
	mux.Handle("/app/", http.StripPrefix("/app/", apiConf.middlewareMetricsInc(http.FileServer(http.Dir(".")))))

	mux.HandleFunc("GET /api/healthz", handleReadinessEndpoint)
	mux.HandleFunc("POST /api/validate_chirp", handlePostChirps)

	mux.HandleFunc("GET /admin/metrics", apiConf.handleMetricsEndpoint)
	mux.HandleFunc("POST /admin/reset", apiConf.handleResetEndpoint)

	serv.ListenAndServe()
}
