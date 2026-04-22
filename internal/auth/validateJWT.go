package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"errors"
)

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