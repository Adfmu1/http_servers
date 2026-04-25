package main

import (
	"github.com/google/uuid"
	"net/http"
)

func handleGetChirps(rw http.ResponseWriter, req *http.Request) {
	query := req.URL.Query().Get("author_id")

	if query != "" {
		usrId, err := uuid.Parse(query)

		if err != nil {
			const errMsg = "User doesnt exist"
			respondWithError(rw, 404, errMsg)
			return
		}

		chirps, err := apiConf.Database.GetChirpsByUserID(req.Context(), usrId)
		
		if err != nil {
			const errMsg = "Something went wrong"
			respondWithError(rw, 500, errMsg)
			return
		}

		respondWithJson(rw, 200, chirps)
		return
	}

	chirps, err := apiConf.Database.GetChirps(req.Context())

	if err != nil {
		const errMsg = "Something went wrong"
		respondWithError(rw, 500, errMsg)
		return
	}

	respondWithJson(rw, 200, chirps)
}