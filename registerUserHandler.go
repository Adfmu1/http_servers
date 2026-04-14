package main

import (
	"encoding/json"
	"net/http"
	"log"
)

func handleRegisterUser(rw http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Email	string	`json:"email"`
	}
	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		const errMsg = "Wrong email"
		respondWithError(rw, 500, errMsg)
		return
	}

	resp, err := apiConf.Database.CreateUser(req.Context(), params.Email)

	if err != nil {
		const errMsg = "Couldnt create user"
		log.Printf("DEBUG: CreateUser failed with error: %v", err)
		respondWithError(rw, 500, errMsg)
		return
	}

	respondWithJson(rw, 201, resp)
}
