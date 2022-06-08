package helper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"net/url"
	"strings"
)

func Unique(arr []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range arr {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func TransferToParentheses(values url.Values) url.Values {
	newValues := url.Values{}

	for key, value := range values {
		keys := strings.Split(key, ".")

		var newKeys string

		if len(keys) > 1 {
			newKeys = keys[0]
			for _, k := range keys[1:] {
				newKeys += "[" + k + "]"
			}
		} else {
			newKeys = strings.Join(keys, "")
		}
		for _, v := range value {
			newValues.Add(newKeys, v)
		}
	}

	return newValues
}

func RsaEncrypt(message string, publicKey string) string {
	block, _ := pem.Decode([]byte(publicKey))
	pub, _ := x509.ParsePKIXPublicKey(block.Bytes)

	ciphertext, err := rsa.EncryptPKCS1v15(
		rand.Reader,
		pub.(*rsa.PublicKey),
		[]byte(message),
	)

	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(ciphertext)
}
