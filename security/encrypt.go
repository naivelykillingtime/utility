package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func hashSecretKey(raw []byte) []byte {
	s := sha256.Sum256(raw)
	return s[:]
}

func encryptAES(raw []byte, secretKey []byte) []byte {
	c, err := aes.NewCipher(secretKey)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	return gcm.Seal(nonce, nonce, raw, nil)
}

func decryptAES(encrypted []byte, secretKey []byte) []byte {
	c, err := aes.NewCipher(secretKey)
	if err != nil {
		panic(err.Error())
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := gcm.NonceSize()
	if len(encrypted) < nonceSize {
		fmt.Println(err)
	}

	nonce, cipherText := encrypted[:nonceSize], encrypted[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		fmt.Println(err)
	}

	return plaintext
}

func encryptString(raw string, secretKey []byte) []byte {
	return encryptAES([]byte(raw), secretKey)
}

func decryptString(encryptedHex string, secretKey []byte) string {
	encrypted, err := hex.DecodeString(encryptedHex)
	if err != nil {
		fmt.Println("wrong hex")
		panic(err)
	}

	return string(decryptAES(encrypted, secretKey))
}
