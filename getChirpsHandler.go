package main

import (
	"net/http"
)

func handleGetChirps(rw http.ResponseWriter, req *http.Request) {
	chirps, err := apiConf.Database.GetChirps(req.Context())

	if err != nil {
		const errMsg = "Something went wrong"
		respondWithError(rw, 500, errMsg)
		return
	}

	respondWithJson(rw, 200, chirps)
}