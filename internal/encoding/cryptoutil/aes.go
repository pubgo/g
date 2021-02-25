package cryptoutil

import (
	"crypto/aes"
	"crypto/cipher"
	"github.com/pubgo/xerror"
)

//1.电码本模式（Electronic Codebook Book (ECB)）；
// 2.密码分组链接模式（Cipher Block Chaining (CBC)）；
// 3.计算器模式（Counter (CTR)）；
// 4.密码反馈模式（Cipher FeedBack (CFB)）；
// 5.输出反馈模式（Output FeedBack (OFB)）

const (
	ECB = iota + 1
	CBC = iota
	CTR = iota
	CFB = iota
	OFB = iota
)

func AesCtrEncrypt(plainText, key []byte, ivAes ...byte) (cipherText []byte, err error) {
	defer xerror.RespErr(&err)

	xerror.Assert(len(key) != 16 && len(key) != 24 && len(key) != 32, ErrTag.ErrKeyLengthSixteen.Error())

	block := xerror.PanicErr(aes.NewCipher(key)).(cipher.Block)

	var iv []byte
	if len(ivAes) != 0 {
		xerror.Assert(len(ivAes) != 16, ErrTag.ErrIvAes.Error())
		iv = ivAes
	} else {
		iv = []byte(ivaes)
	}

	stream := cipher.NewCTR(block, iv)
	cipherText = make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return
}

func AesCtrDecrypt(cipherText, key []byte, ivAes ...byte) (plainText []byte, err error) {
	defer xerror.RespErr(&err)
	xerror.Assert(len(key) != 16 && len(key) != 24 && len(key) != 32, ErrTag.ErrKeyLengthSixteen.Error())

	block := xerror.PanicErr(aes.NewCipher(key)).(cipher.Block)

	var iv []byte
	if len(ivAes) != 0 {
		xerror.Assert(len(ivAes) != 16, ErrTag.ErrIvAes.Error())
		iv = ivAes
	} else {
		iv = []byte(ivaes)
	}

	stream := cipher.NewCTR(block, iv)
	plainText = make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	return
}

// encrypt
func AesCbcEncrypt(plainText, key []byte, ivAes ...byte) (cipherText []byte, err error) {
	defer xerror.RespErr(&err)
	xerror.Assert(len(key) != 16 && len(key) != 24 && len(key) != 32, ErrTag.ErrKeyLengthSixteen.Error())

	block := xerror.PanicErr(aes.NewCipher(key)).(cipher.Block)
	paddingText := _PKCS5Padding(plainText, block.BlockSize())

	var iv []byte
	if len(ivAes) != 0 {
		xerror.Assert(len(ivAes) != 16, ErrTag.ErrIvAes.Error())
		iv = ivAes
	} else {
		iv = []byte(ivaes)
	}

	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText = make([]byte, len(paddingText))
	blockMode.CryptBlocks(cipherText, paddingText)

	return
}

func AesCbcDecrypt(cipherText, key []byte, ivAes ...byte) (plainText []byte, err error) {
	defer xerror.RespErr(&err)
	xerror.Assert(len(key) != 16 && len(key) != 24 && len(key) != 32, ErrTag.ErrKeyLengthSixteen.Error())

	block := xerror.PanicErr(aes.NewCipher(key)).(cipher.Block)

	var iv []byte
	if len(ivAes) != 0 {
		xerror.Assert(len(ivAes) != 16, ErrTag.ErrIvAes.Error())
		iv = ivAes
	} else {
		iv = []byte(ivaes)
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	paddingText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(paddingText, cipherText)
	plainText = xerror.PanicErr(_PKCS5UnPadding(paddingText)).([]byte)

	return
}
