package sign

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
)

func SumSha1Hmac(secret, payload string) string {
	key := []byte(secret)
	hash := hmac.New(sha1.New, key)
	hash.Write([]byte(payload))
	signedBytes := hash.Sum(nil)
	signedString := base64.StdEncoding.EncodeToString(signedBytes)
	return signedString
}

func SumSha256Hmac(secret string, payload ...string) string {
	key := []byte(secret)
	hash := hmac.New(sha256.New, key)
	for _, p := range payload {
		hash.Write([]byte(p))
	}
	signedBytes := hash.Sum(nil)
	signedString := base64.StdEncoding.EncodeToString(signedBytes)
	return signedString
}

func SumSha512Hmac(secret string, payload ...string) string {
	key := []byte(secret)
	hash := hmac.New(sha512.New, key)
	for _, p := range payload {
		hash.Write([]byte(p))
	}
	signedBytes := hash.Sum(nil)
	signedString := base64.StdEncoding.EncodeToString(signedBytes)
	return signedString
}
