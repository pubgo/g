// myxor is a simple encryption algorithm with xor

package cryptoutil

import (
	"encoding/base32"
)

// MyXorEncrypt encrypt
func MyXorEncrypt(text, key []byte) string {
	var _lk = len(key)
	for i := 0; i < len(text); i++ {
		text[i] ^= key[i*i*i%_lk]
	}
	return base32.StdEncoding.EncodeToString(text)
}

//MyXorDecrypt decrypt
func MyXorDecrypt(text string, key []byte) []byte {
	var _lk = len(key)
	_text, _ := base32.StdEncoding.DecodeString(text)
	for i := 0; i < len(_text); i++ {
		_text[i] ^= key[i*i*i%_lk]
	}
	return _text
}
