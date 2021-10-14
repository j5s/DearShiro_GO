package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	_ "fmt"
	"github.com/go-basic/uuid"
)

func GetRememberMe(key []byte, content []byte) (string, error) {
	decodeKey, err := base64.StdEncoding.DecodeString(string(key))
	if err != nil {
		return "", errors.New("decode key error: " + string(key))
	}
	// AES
	block, _ := aes.NewCipher(decodeKey)
	content = padding(content, block.BlockSize())
	iv := []byte(uuid.New())[:16]
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(content))
	blockMode.CryptBlocks(cipherText, content)
	return base64.StdEncoding.EncodeToString(append(iv[:], cipherText[:]...)), nil
}

func padding(plainText []byte, blockSize int) []byte {
	n := blockSize - len(plainText)%blockSize
	temp := bytes.Repeat([]byte{byte(n)}, n)
	plainText = append(plainText, temp...)
	return plainText
}
