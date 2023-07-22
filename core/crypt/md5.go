package crypt

import (
	"crypto/md5"
)

func Md5(data []byte) [16]byte {
	md5New := md5.New()
	md5New.Write(data)
	var bytes = md5New.Sum(nil)
	var bytes16 [16]byte
	copy(bytes16[:16], bytes)
	return bytes16
}
