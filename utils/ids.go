package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateId() string {
	rndBytes := make([]byte, 16)
	rand.Read(rndBytes)
	return hex.EncodeToString(rndBytes)
}
