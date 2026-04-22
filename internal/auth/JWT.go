package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"
	"errors"
	"time"
	"strings"
)

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	claims := &jwt.RegisteredClaims {
		Issuer: "chirpy-access",
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		Subject: userID.String(),
	}
	signMeth := jwt.SigningMethodHS256

	token := jwt.NewWithClaims(signMeth, claims)

	ss, err := token.SignedString([]byte(tokenSecret))

	if err != nil {
		return "", err
	}

	return ss, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	},)

	if err != nil {
		return uuid.Nil, err
	}

	sub, err := token.Claims.GetSubject()

	if err != nil {
		return uuid.Nil, err
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return uuid.Nil, err
	}
	if issuer != string("chirpy-access") {
		return uuid.Nil, errors.New("invalid issuer")
	}

	id, err := uuid.Parse(sub)

	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func GetBearerToken(headers http.Header) (string, error) {
	partsHeader := strings.Fields(headers.Get("Authorization"))

	if len(partsHeader) != 2 || partsHeader[0] != "Bearer" {
		return "", errors.New("malformed authorization header")
	}

	if len(partsHeader[1]) == 0 {
		return "", errors.New("No token in the request")
	}

	return partsHeader[1], nil
}