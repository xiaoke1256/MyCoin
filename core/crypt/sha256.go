package crypt

import (
	"crypto/sha256"
)

func Sha256(data []byte) [32]byte {
	hash := sha256.New()
	hash.Write(data)
	bytes := hash.Sum(nil)
	var bytes32 [32]byte
	copy(bytes32[:32], bytes)
	return bytes32
}

func DoubleSha256(data []byte) [32]byte {
	bytes := Sha256(data)
	copyBytes := make([]byte, len(bytes))
	copy(copyBytes, bytes[:])
	return Sha256(copyBytes)
}
