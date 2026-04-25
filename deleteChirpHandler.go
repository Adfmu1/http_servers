package main

import (
	"github.com/Adfmu1/http_servers/internal/auth"
	"github.com/google/uuid"
	"net/http"
)

func delChirpHandler(rw http.ResponseWriter, req *http.Request) {
	parameter := req.PathValue("chirp_id")
	chirpId, err := uuid.Parse(parameter)

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

	userToken, err := auth.GetBearerToken(req.Header)

	if err != nil {
		const errMsg = "Bad request"
		respondWithError(rw, 401, errMsg)
		return
	}

	usrId, err := auth.ValidateJWT(userToken, apiConf.SecretKey)

	if err != nil {
		const errMsg = "Bad request"
		respondWithError(rw, 401, errMsg)
		return
	}

	if usrId != chirp.UserID {
		const errMsg = "User is not the author of the chirp"
		respondWithError(rw, 403, errMsg)
		return
	}

	err = apiConf.Database.DeleteChirp(req.Context(), chirp.ID)

	if err != nil {
		const errMsg = "Internal error"
		respondWithError(rw, 500, errMsg)
		return
	}

	rw.WriteHeader(204)
}