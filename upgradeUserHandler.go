package main

import (
	"github.com/google/uuid"
	"net/http"
	"encoding/json"
)

func upgradeUserHandler(rw http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Event		string		`json:"event"`
		Data	struct{
			UserID	uuid.UUID	`json:"user_id"`
		}						`json:"data"`
	}
	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		const errMsg = "Wrong request body"
		respondWithError(rw, 400, errMsg)
		return
	}

	if params.Event != "user.upgraded" {
		rw.WriteHeader(204)
		return
	}

	usrId := params.Data.UserID

	err = apiConf.Database.UpgradeUserById(req.Context(), usrId)

	if err != nil {
		const errMsg = "User doesnt exist"
		respondWithError(rw, 404, errMsg)
		return
	}

	rw.WriteHeader(204)
}