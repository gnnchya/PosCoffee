package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"strings"
)

func createHash(key string) string {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	return string(hasher.Sum(nil))
}

func Encrypt(data []byte, passPhrase string) (resultEncode string, err error) {
	block, _ := aes.NewCipher([]byte(createHash(passPhrase)))

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(strings.NewReader(passPhrase), nonce); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return hex.EncodeToString(cipherText), nil
}

func Decrypt(data string, passPhrase string) (resultDecode string, err error) {
	key := []byte(createHash(passPhrase))
	byteData, err := hex.DecodeString(data)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce := byteData[:nonceSize]
	cipherText := byteData[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
