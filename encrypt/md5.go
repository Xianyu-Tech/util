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
