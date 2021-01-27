package cryptoutil

import (
	"crypto/cipher"
	"crypto/des"

	"github.com/pubgo/x/xerror"
)

func DesCbcEncrypt(plainText, key []byte, ivDes ...byte) (cipherText []byte, err error) {
	defer xerror.RespErr(&err)

	xerror.PanicT(len(key) != 8, ErrTag.ErrKeyLengtheEight.Error())
	block := xerror.PanicErr(des.NewCipher(key)).(cipher.Block)
	paddingText := _PKCS5Padding(plainText, block.BlockSize())

	var iv []byte
	if len(ivDes) != 0 {
		xerror.PanicT(len(ivDes) != 8, ErrTag.ErrIvDes.Error())
		iv = ivDes
	} else {
		iv = []byte(ivaes)
	}

	blockMode := cipher.NewCBCEncrypter(block, iv)

	cipherText = make([]byte, len(paddingText))
	blockMode.CryptBlocks(cipherText, paddingText)

	return
}

func DesCbcDecrypt(cipherText, key []byte, ivDes ...byte) (text []byte, err error) {
	defer xerror.RespErr(&err)

	xerror.PanicT(len(key) != 8, ErrTag.ErrKeyLengtheEight.Error())
	block := xerror.PanicErr(des.NewCipher(key)).(cipher.Block)

	var iv []byte
	if len(ivDes) != 0 {
		xerror.PanicT(len(ivDes) != 8, ErrTag.ErrIvDes.Error())
		iv = ivDes
	} else {
		iv = []byte(ivaes)
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	text = xerror.PanicErr(_PKCS5UnPadding(plainText)).([]byte)

	return
}
