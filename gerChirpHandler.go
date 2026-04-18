package main

import (
	"github.com/google/uuid"
	"net/http"
)

func handleGetChirp(rw http.ResponseWriter, req *http.Request) {
	parameter := req.PathValue("chirp_id")
	chirpId ,err:= uuid.Parse(parameter)

	if err != nil {
		const errMsg = "Bad request"
		respondWithError(rw, 400, errMsg)
		return
	}

	chirp, err := apiConf.Database.GetChirp(req.Context(), chirpId)

	if err != nil {
		const errMsg = "Chirp not found"
		respondWithError(rw, 404, errMsg)
		return
	}

	respondWithJson(rw, 200, chirp)
}