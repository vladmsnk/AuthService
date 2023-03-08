package util

import (
	"crypto/sha1"

	"fmt"
)

func HashPassword(password, hashSalt string) string {

	hs := sha1.New()
	hs.Write([]byte(password))
	hs.Write([]byte(hashSalt))

	return fmt.Sprintf("%x", hs.Sum(nil))
}