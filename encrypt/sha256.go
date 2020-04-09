package encryptutil

import (
	"crypto/sha256"
	"fmt"
)

func Sha256Str(data string) string {
	s := sha256.New()

	s.Write([]byte(data))
	resp := fmt.Sprintf("%x", s.Sum(nil))

	return resp
}

func Sha256(data []byte) []byte {
	s := sha256.New()

	s.Write(data)
	resp := fmt.Sprintf("%x", s.Sum(nil))

	return []byte(resp)
}
