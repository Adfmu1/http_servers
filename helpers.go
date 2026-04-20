package main

import (
	"github.com/Adfmu1/http_servers/internal/database"
	"github.com/google/uuid"
	"time"
	"encoding/json"
	"net/http"
	"strings"
	"slices"
)

type UserNoPass struct {
    ID        uuid.UUID     `json:"id"`
    CreatedAt time.Time     `json:"created_at"`
    UpdatedAt time.Time     `json:"updated_at"`
    Email     string        `json:"email"`
}

func respondWithError(rw http.ResponseWriter, code int, msg string) {
	type errorStruct struct {
		ErrMsg	string	`json:"error"`
	}
	errDat := errorStruct{
		ErrMsg:	msg,
	}
	respondWithJson(rw, code, errDat)
}

func respondWithJson(rw http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		const errMsg = "Something went wrong"
		respondWithError(rw, 500, errMsg)
		return
	}
	
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)
	rw.Write(data)
}

func filterProfaneWords(chirp string) string {
	// banned words
	profaneWords := []string{"kerfuffle", "sharbert", "fornax"}

	words := strings.Split(chirp, " ")
	wordsFiltered := make([]string, len(words))
	copy(wordsFiltered, words)

	for i := 0; i < len(words); i++ {
		if slices.Contains(profaneWords, strings.ToLower(words[i])) {
			wordsFiltered[i] = "****"
		}
	}

	return strings.Join(wordsFiltered, " ")
}

func removePasswordFromUser(usr database.User) UserNoPass {
	return UserNoPass{
		ID: usr.ID,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
		Email: usr.Email,
	}
}