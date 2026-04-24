package main

import (
	"github.com/Adfmu1/http_servers/internal/auth"
	"net/http"
	"time"
)

func handleRefresh(rw http.ResponseWriter, req *http.Request) {
	refreshTokenID, err := auth.GetBearerToken(req.Header)

	if err != nil {
		const errMsg = "Malformed request"
		respondWithError(rw, 401, errMsg)
		return
	}

	refreshToken, err := apiConf.Database.GetRefreshTokenInfo(req.Context(), refreshTokenID)

	if err != nil {
		const errMsg = "Authentication failed"
		respondWithError(rw, 401, errMsg)
		return
	}

	now := time.Now()
	if refreshToken.RevokedAt.Valid || now.After(refreshToken.ExpiresAt) {
		const errMsg = "Authentication failed"
		respondWithError(rw, 401, errMsg)
		return
	}

	token, err := auth.MakeJWT(refreshToken.UserID, apiConf.SecretKey, time.Duration(3600) * time.Second)

	if err != nil {
		const errMsg = "Internal error"
		respondWithError(rw, 500, errMsg)
		return
	}

	tokenStruct := struct{
		Token	string	`json:"token"`
	} {
		Token: token,
	}
	
	respondWithJson(rw, 200, tokenStruct)
}