package main

import (
	"github.com/Adfmu1/http_servers/internal/auth"
	"github.com/Adfmu1/http_servers/internal/database"
	"encoding/json"
	"net/http"
)

func editUserDataHandler(rw http.ResponseWriter, req *http.Request) {
	type parameters struct {
		NewPassword	string		`json:"password"`
		NewEmail	string		`json:"email"`
	}
	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		const errMsg = "Wrong request body"
		respondWithError(rw, 400, errMsg)
		return
	}

	accessToken, err := auth.GetBearerToken(req.Header)

	if err != nil {
		const errMsg = "Invalid JWT token"
		respondWithError(rw, 401, errMsg)
		return
	}

	id, err := auth.ValidateJWT(accessToken, apiConf.SecretKey)

	if err != nil {
		const errMsg = "Invalid JWT token"
		respondWithError(rw, 401, errMsg)
		return
	}

	newPasswordHashed, err := auth.HashPassword(params.NewPassword)

	if err != nil {
		const errMsg = "Internal error"
		respondWithError(rw, 500, errMsg)
		return
	}

	newParams := database.UpdateEmailAndPasswordParams{
		ID:	id,
		Email:	params.NewEmail,
		HashedPassword: newPasswordHashed,
	}

	err = apiConf.Database.UpdateEmailAndPassword(req.Context(), newParams)

	if err != nil {
		const errMsg = "Database Conflict"
		respondWithError(rw, 409, errMsg)
		return
	}

	newUserData, err := apiConf.Database.GetUser(req.Context(), newParams.Email)

	if err != nil {
		const errMsg = "Internal error"
		respondWithError(rw, 500, errMsg)
		return
	}

	respondWithJson(rw, 200, removePasswordFromUser(newUserData))
}