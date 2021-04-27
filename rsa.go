package WeWorkFinanceSDK

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

func RSADecrypt(privateKey string, ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("RSAPrivateKey format invalid, decode failed")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func RSADecryptBase64(privateKey string, cryptoText string) ([]byte, error) {
	encryptedData, _ := base64.StdEncoding.DecodeString(cryptoText)
	return RSADecrypt(privateKey, encryptedData)
}
