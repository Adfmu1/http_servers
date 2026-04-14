package main

import (
	"net/http"
)

func (cfg *apiConfig) handleResetEndpoint(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(200)
	cfg.fileserverHits.Store(0)
}