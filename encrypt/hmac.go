package encryptutil

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

func HmacStr(key string, data string) string {
	h := hmac.New(md5.New, []byte(key))

	h.Write([]byte(data))
	resp := hex.EncodeToString(h.Sum(nil))

	return resp
}

func Hmac(key string, data []byte) []byte {
	h := hmac.New(md5.New, []byte(key))

	h.Write(data)
	resp := hex.EncodeToString(h.Sum(nil))

	return []byte(resp)
}
