package project

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter the password for verification")
	password, _ := reader.ReadString('\n')
	password = password[:len(password)-1] // Remove the newline character
	hashed := HashPassword(password)
	return VerifyPassword(hashed, password)
}
