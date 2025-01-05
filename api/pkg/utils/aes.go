package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"github.com/spf13/viper"
)

func AESEncrypt(plainText string) (string, error) {
	key := []byte(viper.GetString("application.key"))
	if len(key) != 32 {
		return "", errors.New("key must be 32 bytes long")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func AESDecrypt(cipherText string) (string, error) {
	key := []byte(viper.GetString("application.key"))
	if len(key) != 32 {
		return "", errors.New("key must be 32 bytes long")
	}

	encryptedData, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(encryptedData) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, cipherData := encryptedData[:nonceSize], encryptedData[nonceSize:]

	plainText, err := aesGCM.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
