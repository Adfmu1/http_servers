package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"slices"
)

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
