package auth

import (
	"net/http"
	"strings"
	"errors"
)

func GetAPIKey(headers http.Header) (string, error) {
	partsHeader := strings.Fields(headers.Get("Authorization"))

	if len(partsHeader) != 2 || partsHeader[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return partsHeader[1], nil
}