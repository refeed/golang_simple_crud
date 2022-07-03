package auth

import (
	"crypto/sha256"
	"fmt"
)

// Convert a password string to a SHA-256 hash string
func HashPassword(password string) string {
	hashByte := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hashByte)
}
