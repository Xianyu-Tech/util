package encryptutil

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
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

	return publicKey.(*rsa.PublicKey), nil
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

	return privateKey, nil
}

// RSA加密
func RsaEncrypt(data []byte, publickey *rsa.PublicKey) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, publickey, data)
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

func RsaSignWithSha256ToStr(data []byte, privatekey *rsa.PrivateKey) (string, error) {
	sign, err := RsaSignWithSha256(data, privatekey)

	if err != nil {
		return "", err
	}

	signStr := hex.EncodeToString(sign)

	return signStr, nil
}

func RsaSignWithSha256ToBase64(data []byte, privatekey *rsa.PrivateKey) (string, error) {
	sign, err := RsaSignWithSha256(data, privatekey)

	if err != nil {
		return "", err
	}

	signStr := base64.StdEncoding.EncodeToString(sign)

	return signStr, nil
}

// RSA验签
func RsaVerifyWithSha256(data, sign []byte, publickey *rsa.PublicKey) bool {
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

func RsaVerifyWithSha256FromStr(data []byte, sign string, publickey *rsa.PublicKey) bool {
	signBytes, err := hex.DecodeString(sign)

	if err != nil {
		return false
	}

	resp := RsaVerifyWithSha256(data, signBytes, publickey)

	return resp
}
