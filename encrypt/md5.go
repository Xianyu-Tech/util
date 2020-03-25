package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5String(input []byte) (output string) {
	m5 := md5.New()
	m5.Write(input)
	output = hex.EncodeToString(m5.Sum(nil))
	return output
}

func Md5(input []byte) (output []byte) {
	m5 := md5.New()
	m5.Write(input)
	output = m5.Sum(nil)
	return output
}
