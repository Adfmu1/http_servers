package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
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