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

func HmacWithSha256ToStr(key []byte, data []byte) string {
	resp := HmacWithSha256(key, data)

	respStr := hex.EncodeToString(resp)

	return respStr
}
