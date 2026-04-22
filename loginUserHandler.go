	package main

	import (
		"github.com/Adfmu1/http_servers/internal/auth"
		"github.com/google/uuid"
		"net/http"
		"encoding/json"
		"time"
	)

	func handleLoginUser (rw http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		type loginParams struct {
			ExpiresInSeconds	*int	`json:"expires_in_seconds"`
			Email          		string	`json:"email"`
			HashedPassword 		string	`json:"password"`
		}
		params := loginParams{}
		err := decoder.Decode(&params)

		if err != nil {
			const errMsg = "Wrong request"
			respondWithError(rw, 400, errMsg)
			return
		}

		if params.ExpiresInSeconds == nil  {
			expTime := 3600
			params.ExpiresInSeconds = &expTime
		}else if *params.ExpiresInSeconds > 3600 || *params.ExpiresInSeconds <= 0 {
			*params.ExpiresInSeconds = 3600
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

		token, err := auth.MakeJWT(dbUsr.ID, apiConf.SecretKey, time.Duration(*params.ExpiresInSeconds) * time.Second)

		if err != nil {
			const errMsg = "Could not create web token"
			respondWithError(rw, 500, errMsg)
			return			
		}

		type respStruct struct {
			ID        	uuid.UUID     	`json:"id"`
			CreatedAt 	time.Time     	`json:"created_at"`
			UpdatedAt 	time.Time     	`json:"updated_at"`
			Email     	string        	`json:"email"`
			Token		string			`json:"token"`
		}

		respUsr := respStruct{
			ID: dbUsr.ID,
			CreatedAt: dbUsr.CreatedAt,
			UpdatedAt: dbUsr.UpdatedAt,
			Email: dbUsr.Email,
			Token: token,
		}

		respondWithJson(rw, 200, respUsr)
	}