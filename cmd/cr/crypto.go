package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"crypto/rand"
	"io"
)

func mdHashing(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	return hex.EncodeToString(md5Hash[:]) // by referring to it as a string
}

func encrypt(value []byte, keyPhrase string) []byte {

	aesBlock, err := aes.NewCipher([]byte(mdHashing(keyPhrase)))
	if err != nil {
		panic(err)
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcmInstance.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	cipheredText := gcmInstance.Seal(nonce, nonce, value, nil)

	return cipheredText
}

func decrypt(ciphered []byte, keyPhrase string) []byte {
	hashedPhrase := mdHashing(keyPhrase)
	aesBlock, err := aes.NewCipher([]byte(hashedPhrase))
	if err != nil {
		panic(err)
	}
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		panic(err)
	}
	nonceSize := gcmInstance.NonceSize()
	nonce, cipheredText := ciphered[:nonceSize], ciphered[nonceSize:]
	originalText, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		panic(err)
	}
	return originalText
}
