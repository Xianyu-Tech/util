package randutil

import (
	"crypto/rand"
	"encoding/base64"
)

func RandBytes(n int) ([]byte, error) {
	data := make([]byte, n)

	_, err := rand.Read(data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func RandBase64Str(n int) (string, error) {
	data, err := RandBytes(n)

	return base64.URLEncoding.EncodeToString(data), err
}
