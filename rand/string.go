package randutil

import (
	"math/rand"
)

func RandCharStr(n int) string {
	bytesInit := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	data := make([]byte, 0)

	for i := 0; i < n; i++ {
		data = append(data, bytesInit[rand.Intn(len(bytesInit))])
	}

	return string(data)
}

func RandNumStr(n int) string {
	bytesInit := []byte("0123456789")

	data := make([]byte, 0)

	for i := 0; i < n; i++ {
		data = append(data, bytesInit[rand.Intn(len(bytesInit))])
	}

	return string(data)
}
