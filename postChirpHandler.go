package main

import (
	"github.com/Adfmu1/http_servers/internal/database"
	"encoding/json"
	"net/http"
	"github.com/google/uuid"
)

func handlePostChirps(rw http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Body	string		`json:"body"`
		UserID	uuid.UUID	`json:"user_id"`
	}
	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	// ======== CHECK REQUEST DATA ========
	// any error occurs
	if err != nil {
		const errMsg = "Something went wrong"
		respondWithError(rw, 500, errMsg)
		return
	}
	// length of chirp too long
	if len(params.Body) > 140 {
		const errMsg = "Chirp is too long"
		respondWithError(rw, 400, errMsg)
		return
	}
	params.Body = filterProfaneWords(params.Body)
	// request ok
	chirp, err := apiConf.Database.PostChirp(
		req.Context(), database.PostChirpParams{
			Body: params.Body,
			UserID: params.UserID,
		})

	if err != nil {
		const errMsg = "Something went wrong while creating Chirp"
		respondWithError(rw, 500, errMsg)
		return
	}

	respondWithJson(rw, 201, chirp)
}
