package auth

import (
	"fmt"
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
	password1 := "helloworldpasssword"
	password2 := "helloworldpassswordabc"
	hash1, _ := HashPassword(password1)
	hash2, _ := HashPassword(password2)

	tests := []struct {
		name     string
		password string
		hash     string
		wantErr  bool
	}{{
		name:     "Correct password",
		password: password1,
		hash:     hash1,
		wantErr:  false,
	},
		{
			name:     "Incorrect password",
			password: "wrongPassword",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Password doesn't match different hash",
			password: password1,
			hash:     hash2,
			wantErr:  true,
		},
		{
			name:     "Empty password",
			password: "",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Invalid hash",
			password: password1,
			hash:     "invalidhash",
			wantErr:  true,
		},
		{
			name:     "Correct password",
			password: password2,
			hash:     hash2,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckPasswordHash(tt.password, tt.hash); (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
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
	userID := uuid.New()
	validToken, _ := MakeJWT(userID, "secret", time.Hour)

	tests := []struct {
		name        string
		tokenString string
		tokenSecret string
		wantUserID  uuid.UUID
		wantErr     bool
	}{
		{
			name:        "Valid token",
			tokenString: validToken,
			tokenSecret: "secret",
			wantUserID:  userID,
			wantErr:     false,
		},
		{
			name:        "Invalid token",
			tokenString: "invalid.token.string",
			tokenSecret: "secret",
			wantUserID:  uuid.Nil,
			wantErr:     true,
		},
		{
			name:        "Wrong secret",
			tokenString: validToken,
			tokenSecret: "wrong_secret",
			wantUserID:  uuid.Nil,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserID, err := ValidateJWT(tt.tokenString, tt.tokenSecret)
			fmt.Println(gotUserID)
			fmt.Println(tt.wantUserID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUserID != tt.wantUserID {
				t.Errorf("ValidateJWT() gotUserID = %v, want %v", gotUserID, tt.wantUserID)
			}
		})
	}
}
