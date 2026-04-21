package auth

import (
	"net/http"
	"strings"
	"errors"
)

func GetBearerToken(headers http.Header) (string, error) {
	partsHeader := strings.Split(headers.Get("Authorization"), " ")

	if len(partsHeader) != 2 || partsHeader[0] != "Bearer" {
		return "", errors.New("malformed authorization header")
	}

	if len(partsHeader[1]) == 0 {
		return "", errors.New("No token in the request")
	}

	return partsHeader[1], nil
}