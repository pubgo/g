package cryptoutil

import (
	"bytes"
)

func _PKCS5Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - len(plainText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(plainText, padText...)
	return newText
}

func _PKCS5UnPadding(plainText []byte) ([]byte, error) {
	length := len(plainText)
	number := int(plainText[length-1])
	if number >= length {
		return nil, ErrTag.ErrPaddingSize
	}
	return plainText[:length-number], nil
}
