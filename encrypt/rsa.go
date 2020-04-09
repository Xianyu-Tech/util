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

// RSA加密
func RsaEncrypt(data []byte, publickey []byte) ([]byte, error) {
	block, _ := pem.Decode(publickey)

	if block == nil {
		return nil, errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	pub := pubInterface.(*rsa.PublicKey)

	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

// RSA解密
func RsaDecrypt(ciphertext []byte, privatekey []byte) ([]byte, error) {
	block, _ := pem.Decode(privatekey)

	if block == nil {
		return nil, errors.New("private key error!")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
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
