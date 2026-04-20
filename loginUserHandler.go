package main

import (
	"github.com/Adfmu1/http_servers/internal/database"
	"github.com/Adfmu1/http_servers/internal/auth"
	"net/http"
	"encoding/json"
)

func handleLoginUser (rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	params := database.CreateUserParams{}
	err := decoder.Decode(&params)

	if err != nil {
		const errMsg = "Wrong request"
		respondWithError(rw, 500, errMsg)
		return
	}

	dbUsr, err := apiConf.Database.GetUser(req.Context(), params.Email)

	if err != nil {
		const errMsg = "Incorrect email or password"
		respondWithError(rw, 401, errMsg)
		return
	}

	passOk, err := auth.CheckPasswordHash(params.HashedPassword, dbUsr.HashedPassword)

	if err != nil || !passOk {
		const errMsg = "Incorrect email or password"
		respondWithError(rw, 401, errMsg)
		return
	}

	respondWithJson(rw, 200, removePasswordFromUser(dbUsr))	
}