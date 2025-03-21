package common

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(plaintext string) string {
	hash := sha256.Sum256([]byte(plaintext))
	return hex.EncodeToString(hash[:])
}
