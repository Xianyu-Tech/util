package encryptutil

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HmacWithSha256(key []byte, data []byte) []byte {
	h := hmac.New(sha256.New, key)

	h.Write(data)

	resp := h.Sum(nil)

	return resp
}
