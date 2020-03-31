package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func GenerateHash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}

func ValidateHash(input string, hash string) bool {
	to_check := GenerateHash(input)
	if to_check == hash {
		return true
	} else {
		return false
	}
}

func GenerateToken() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
