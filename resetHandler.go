package main

import (
	"net/http"
)

func (cfg *apiConfig) handleResetEndpoint(rw http.ResponseWriter, req *http.Request) {
	if cfg.Platform != "dev" {
		const errMsg = "403 Forbidden"
		respondWithError(rw, 403, errMsg)
		return
	}

	err := apiConf.Database.DeleteUser(req.Context())

	if err != nil {
		const errMsg = "Couldnt delete users"
		respondWithError(rw, 500, errMsg)
		return
	}

	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	rw.WriteHeader(200)
	cfg.fileserverHits.Store(0)
}