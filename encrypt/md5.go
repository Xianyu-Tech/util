package encryptutil

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(data []byte) []byte {
	m := md5.New()

	m.Write(data)

	resp := m.Sum(nil)

	return resp
}

func Md5ToStr(data []byte) string {
	resp := Md5(data)

	respStr := hex.EncodeToString(resp)

	return respStr
}
