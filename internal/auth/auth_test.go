package auth

import "testing"

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
