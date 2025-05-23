package helper

import (
	"crypto/aes"
	cpr "crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	hinf "restuwahyu13/shopping-cart/internal/domain/interface/helper"

	"golang.org/x/crypto/scrypt"
)

type crypto struct{}

func NewCrypto() hinf.ICrypto {
	return &crypto{}
}

func (h *crypto) AES256Encrypt(secretKey, plainText string) (string, error) {
	secretKeyByte := make([]byte, len(secretKey))
	secretKeyByte = []byte(secretKey)

	plainTextByte := make([]byte, len(plainText))
	plainTextByte = []byte(plainText)

	tagSize := 16

	if len(secretKeyByte) < 32 {
		return "", errors.New("Secretkey length mismatch")
	}

	key, err := scrypt.Key([]byte(secretKey), []byte("salt"), 1024, 8, 1, 32)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cpr.NewGCMWithTagSize(block, tagSize)
	if err != nil {
		return "", err
	}

	nonceSize := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonceSize); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonceSize, nonceSize, []byte(plainTextByte), nil)

	return hex.EncodeToString(cipherText), nil
}

func (h *crypto) AES256Decrypt(secretKey string, cipherText string) (string, error) {
	secretKeyByte := make([]byte, len(secretKey))
	secretKeyByte = []byte(secretKey)
	tagSize := 16

	if len(secretKeyByte) < 32 {
		return "", errors.New("Secretkey length mismatch")
	}

	key, err := scrypt.Key(secretKeyByte, []byte("salt"), 1024, 8, 1, 32)
	if err != nil {
		return "", err
	}

	cipherTextByte, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cpr.NewGCMWithTagSize(block, tagSize)
	if err != nil {
		return "", err
	}

	nonceSize := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonceSize); err != nil {
		return "", err
	} else if len(cipherTextByte) < len(nonceSize) {
		return "", errors.New("Cipher text to short")
	}

	nonce, ciphertext := cipherTextByte[:len(nonceSize)], cipherTextByte[len(nonceSize):]
	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func (h *crypto) HMACSHA512Sign(secretKey, data string) (string, error) {
	hashHMAC512 := hmac.New(sha512.New, []byte(secretKey))

	if _, err := hashHMAC512.Write([]byte(data)); err != nil {
		return "", err
	}

	return hex.EncodeToString(hashHMAC512.Sum(nil)), nil
}

func (h *crypto) HMACSHA512Verify(secretKey, data, hash string) bool {
	hashHMAC512 := hmac.New(sha512.New, []byte(secretKey))

	if _, err := hashHMAC512.Write([]byte(data)); err != nil {
		return false
	}

	return hmac.Equal([]byte(hash), hashHMAC512.Sum(nil))
}

func (h *crypto) SHA256Sign(plainText string) (string, error) {
	hashSHA256 := sha256.New()

	if _, err := hashSHA256.Write([]byte(plainText)); err != nil {
		return "", err
	}
	return hex.EncodeToString(hashSHA256.Sum(nil)), nil
}

func (h *crypto) SHA512Sign(plainText string) (string, error) {
	hashSHA512 := sha512.New()

	if _, err := hashSHA512.Write([]byte(plainText)); err != nil {
		return "", err
	}
	return hex.EncodeToString(hashSHA512.Sum(nil)), nil
}
