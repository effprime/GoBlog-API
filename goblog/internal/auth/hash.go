package auth

import (
	"crypto/sha1"
	"encoding/hex"
)

func GenerateHash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func validateHash(input string, hash string) bool {
	if GenerateHash(input) == hash {
		return true
	}
	return false
}
