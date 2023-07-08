package crypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func Sha256(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	bytes := hash.Sum(nil)
	return bytes
}
