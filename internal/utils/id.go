package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func CreateID(args ...int) (string, error) {
	length := 16 // Default length

	if len(args) > 0 {
		length = args[0]
	}

	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
