package aes

import (
	cryptoAes "crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
)

func Decrypt(cipherText string, password string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	if string(data[:8]) != "Salted__" {
		return "", errors.New("invalid crypto js aes encryption")
	}

	salt := data[8:16]
	cipherBytes := data[16:]
	key, iv, err := DefaultEvpKDF([]byte(password), salt)
	if err != nil {
		return "", err
	}

	block, err := cryptoAes.NewCipher(key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherBytes, cipherBytes)

	result := PKCS5UnPadding(cipherBytes)
	return string(result), nil
}

func EncryptWithSalt(content string, password string, salt []byte) (string, error) {
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	key, iv, err := DefaultEvpKDF([]byte(password), salt)

	block, err := cryptoAes.NewCipher(key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	cipherBytes := PKCS5Padding([]byte(content), cryptoAes.BlockSize)
	mode.CryptBlocks(cipherBytes, cipherBytes)

	data := make([]byte, 16+len(cipherBytes))
	copy(data[:8], []byte("Salted__"))
	copy(data[8:16], salt)
	copy(data[16:], cipherBytes)

	cipherText := base64.StdEncoding.EncodeToString(data)
	return cipherText, nil
}

func Encrypt(content string, password string) (string, error) {
	salt := make([]byte, 8)
	return EncryptWithSalt(content, password, salt)
}
