package main

import (
	"github.com/Adfmu1/http_servers/internal/database"
	"github.com/Adfmu1/http_servers/internal/auth"
	"encoding/json"
	"net/http"
)

func handlePostChirps(rw http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Body	string		`json:"body"`
	}
	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		const errMsg = "Wrong request body"
		respondWithError(rw, 400, errMsg)
		return
	}

	usrToken, err := auth.GetBearerToken(req.Header)

	if err != nil {
		const errMsg = "Invalid JWT token"
		respondWithError(rw, 401, errMsg)
		return
	}

	id, err := auth.ValidateJWT(usrToken, apiConf.SecretKey)

	if err != nil {
		const errMsg = "Invalid JWT token"
		respondWithError(rw, 401, errMsg)
		return
	}
	
	// ======== CHECK REQUEST DATA ========
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
			UserID: id,
		})

	if err != nil {
		const errMsg = "Something went wrong while creating Chirp"
		respondWithError(rw, 500, errMsg)
		return
	}

	respondWithJson(rw, 201, chirp)
}
