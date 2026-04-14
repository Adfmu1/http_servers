package main

import (
	"encoding/json"
	"net/http"
)

func handlePostChirps(rw http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Body	string	`json:"body"`
	}
	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	params.Body = filterProfaneWords(params.Body)

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
	// response ok
	type response struct {
		Body	string	`json:"cleaned_body"`
	}
	resp := response{
		Body:	params.Body,
	}
	respondWithJson(rw, 200, resp)
}
