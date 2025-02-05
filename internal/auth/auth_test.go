package auth

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

const tokenSecret = "supersecret"

func TestHashPassword(t *testing.T) {
	password := "<PASSWORD>"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Error("Hashing password failed: %w", err)
	}

	err = CheckPasswordHash(password, hashedPassword)
	if err != nil {
		t.Error("Checking password failed: %w", err)
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "helloworldpasssword"
	wrongPassword := "<PASSWORD>"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Error("Hashing password failed: %w", err)
	}

	err = CheckPasswordHash(wrongPassword, hashedPassword)
	if err == nil {
		t.Error("Checking password failed: %w", err)
	}
}

func TestMakeJWT(t *testing.T) {
	userId := uuid.New()
	expiresTime := 60 * time.Second

	_, err := MakeJWT(userId, tokenSecret, expiresTime)
	if err != nil {
		t.Errorf("MakeJWT failed: %v", err)
	}
}

func TestValidateJWT(t *testing.T) {
	userId := uuid.New()
	expiresTime := 60 * time.Second

	jwtToken, err := MakeJWT(userId, tokenSecret, expiresTime)
	if err != nil {
		t.Error("MakeJWT failed: %w", err)
	}

	userUUID, err := ValidateJWT(jwtToken, tokenSecret)
	if err != nil {
		t.Error("ValidateJWT failed: %w", err)
	}

	if userUUID != userId {
		t.Error("ValidateJWT failed: userID does not match")
	}
}
