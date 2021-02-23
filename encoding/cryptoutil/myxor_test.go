package cryptoutil

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestMyXorEncrypt(t *testing.T) {
	fmt.Println(base64.StdEncoding.EncodeToString(MyXorEncrypt([]byte("hello"), []byte("123456"))))
}

func TestMyXorDecrypt(t *testing.T) {
	_e := MyXorEncrypt([]byte("sxscssbahello6sxnebwebfebhbhsbhbhbhbhbhb"), []byte("123456"))
	fmt.Println(string(_e))
	fmt.Println(string(MyXorDecrypt(_e, []byte("123456"))))
}
