package main

import (
	"github.com/google/uuid"
	"net/http"
	"sort"
)

func handleGetChirps(rw http.ResponseWriter, req *http.Request) {
	query := req.URL.Query().Get("author_id")
	sortType := req.URL.Query().Get("sort")

	if query != "" {
		usrId, err := uuid.Parse(query)

		if err != nil {
			const errMsg = "User doesnt exist"
			respondWithError(rw, 404, errMsg)
			return
		}

		chirps, err := apiConf.Database.GetChirpsByUserID(req.Context(), usrId)
		
		if err != nil {
			const errMsg = "Something went wrong"
			respondWithError(rw, 500, errMsg)
			return
		}

		if sortType == "desc" {
			sort.Slice(chirps, func(i, j int) bool {return chirps[i].CreatedAt.After(chirps[j].CreatedAt)} )
		}

		respondWithJson(rw, 200, chirps)
		return
	}

	chirps, err := apiConf.Database.GetChirps(req.Context())

	if err != nil {
		const errMsg = "Something went wrong"
		respondWithError(rw, 500, errMsg)
		return
	}

	if sortType == "desc" {
		sort.Slice(chirps, func(i, j int) bool {return chirps[i].CreatedAt.After(chirps[j].CreatedAt)} )
	}

	respondWithJson(rw, 200, chirps)
}