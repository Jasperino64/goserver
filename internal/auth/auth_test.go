package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "testPassword"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if hashedPassword == "" {
		t.Fatal("Hashed password should not be empty")
	}
}
func TestCheckPasswordHash(t *testing.T) {
	password := "testPassword"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	err = CheckPasswordHash(password, hashedPassword)
	if err != nil {
		t.Fatalf("Password check failed: %v", err)
	}

	wrongPassword := "wrongPassword"
	err = CheckPasswordHash(wrongPassword, hashedPassword)
	if err == nil {
		t.Fatal("Expected password check to fail for wrong password")
	}
}