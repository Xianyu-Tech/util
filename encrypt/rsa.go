package encryptutil

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func GenRsaPublicKey(data []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(data)

	if block == nil {
		return nil, errors.New("data illegal")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	return publicKey.(*rsa.PublicKey)

}

func GenRsaPrivateKey(data []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(data)

	if block == nil {
		return nil, errors.New("data illegal")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	return privateKey
}

// RSA加密
func RsaEncrypt(data []byte, publickey *rsa.PublicKey) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
}

// RSA解密
func RsaDecrypt(data []byte, privatekey *rsa.PrivateKey) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privatekey, data)
}

// RSA签名
func RsaSignWithSha256(data []byte, privatekey *rsa.PrivateKey) ([]byte, error) {
	h := sha256.New()
	h.Write(data)

	bytes := h.Sum(nil)

	sign, err := rsa.SignPKCS1v15(rand.Reader, privatekey, crypto.SHA256, bytes)

	if err != nil {
		return nil, err
	}

	return sign, nil
}

// RSA验签
func RsaVerySignWithSha256(data, sign []byte, publickey *rsa.PublicKey) bool {
	h := sha256.New()
	h.Write(data)

	bytes := h.Sum(nil)

	//验证数字签名
	err := rsa.VerifyPKCS1v15(publickey, crypto.SHA256, bytes, sign)

	if err == nil {
		return true
	}

	return false
}
