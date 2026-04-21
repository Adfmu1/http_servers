// auth_test.go
package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	userID := uuid.New()
	secret := "my-secret-key"

	// Create token
	token, err := MakeJWT(userID, secret, time.Hour)
	if err != nil {
		t.Fatalf("unexpected error creating token: %v", err)
	}

	// Validate token
	gotID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("unexpected error validating token: %v", err)
	}

	// Check returned user ID
	if gotID != userID {
		t.Fatalf("got %v, want %v", gotID, userID)
	}
}

func TestExpiredJWTRejected(t *testing.T) {
	userID := uuid.New()
	secret := "my-secret-key"

	// Expired token (already expired 1 minute ago)
	token, err := MakeJWT(userID, secret, -time.Minute)
	if err != nil {
		t.Fatalf("unexpected error creating token: %v", err)
	}

	_, err = ValidateJWT(token, secret)
	if err == nil {
		t.Fatal("expected error for expired token, got nil")
	}
}

func TestJWTWrongSecretRejected(t *testing.T) {
	userID := uuid.New()

	token, err := MakeJWT(userID, "correct-secret", time.Hour)
	if err != nil {
		t.Fatalf("unexpected error creating token: %v", err)
	}

	_, err = ValidateJWT(token, "wrong-secret")
	if err == nil {
		t.Fatal("expected error for wrong secret, got nil")
	}
}