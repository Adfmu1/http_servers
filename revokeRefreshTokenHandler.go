package main

import (
	"github.com/Adfmu1/http_servers/internal/auth"
	"net/http"
)

func handleRevoke(rw http.ResponseWriter, req *http.Request) {
	refreshTokenID, err := auth.GetBearerToken(req.Header)

	if err != nil {
		const errMsg = "Malformed request"
		respondWithError(rw, 401, errMsg)
		return
	}

	err = apiConf.Database.RevokeRefreshToken(req.Context(), refreshTokenID)

	if err != nil {
		const errMsg = "Internal error"
		respondWithError(rw, 500, errMsg)
		return
	}

	rw.WriteHeader(204)
}