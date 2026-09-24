package project

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func VerifyPassword(hashedPassword, password string) bool {
	return HashPassword(password) == hashedPassword
}

func TestPasswordVerification() bool {
	password := "examplePassword"
	hashed := HashPassword(password)
	return VerifyPassword(hashed, password)
}
