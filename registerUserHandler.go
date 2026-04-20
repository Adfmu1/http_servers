package main

import (
	"github.com/Adfmu1/http_servers/internal/database"
	"github.com/Adfmu1/http_servers/internal/auth"
	"encoding/json"	
	"net/http"
)

func handleRegisterUser(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	params := database.CreateUserParams{}
	err := decoder.Decode(&params)

	if err != nil {
		const errMsg = "Wrong email"
		respondWithError(rw, 500, errMsg)
		return
	}

	params.HashedPassword, err = auth.HashPassword(params.HashedPassword)

	if err != nil {
		const errMsg = "Couldnt hash user password"
		respondWithError(rw, 500, errMsg)
		return
	}

	resp, err := apiConf.Database.CreateUser(req.Context(), params)

	if err != nil {
		const errMsg = "Couldnt create user"
		respondWithError(rw, 500, errMsg)
		return
	}

	respondWithJson(rw, 201, removePasswordFromUser(resp))
}
