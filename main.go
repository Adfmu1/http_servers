package main

import (
	"net/http"
	"sync/atomic"
)

// struct to hold number of requests to the server
type apiConfig struct {
	fileserverHits atomic.Int32
}

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
	// readiness endpoint handler
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

// middleware method for counting requests
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}
