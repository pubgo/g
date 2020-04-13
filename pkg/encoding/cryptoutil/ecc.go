package cryptoutil

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"github.com/pubgo/g/pkg/encoding/hashutil"
	"github.com/pubgo/g/xerror"
	"math/big"
)

func EccSign(msg []byte, Key []byte) (r []byte, s []byte, err error) {
	defer xerror.RespErr(&err)

	block, _ := pem.Decode(Key)
	privateKey := xerror.PanicErr(x509.ParseECPrivateKey(block.Bytes)).(*ecdsa.PrivateKey)

	ir, is, _err := ecdsa.Sign(rand.Reader, privateKey, hashutil.Sha256(msg))
	xerror.Panic(_err)

	r = xerror.PanicErr(ir.MarshalText()).([]byte)
	s = xerror.PanicErr(is.MarshalText()).([]byte)

	return
}

func EccVerifySign(msg []byte, Key []byte, rText, sText []byte) (b bool, err error) {
	defer xerror.RespErr(&err)

	block, _ := pem.Decode(Key)

	publicKey := xerror.PanicErr(x509.ParsePKIXPublicKey(block.Bytes)).(*ecdsa.PublicKey)

	var r, s big.Int
	xerror.Panic(r.UnmarshalText(rText))
	xerror.Panic(s.UnmarshalText(sText))
	b = ecdsa.Verify(publicKey, hashutil.Sha256(msg), &r, &s)

	return
}

// The public key and plaintext are passed in for encryption
func EccEncrypt(plainText, key []byte) (cryptText []byte, err error) {
	defer xerror.RespErr(&err)

	block, _ := pem.Decode(key)

	publicKey := xerror.PanicErr(x509.ParsePKIXPublicKey(block.Bytes)).(*ecdsa.PublicKey)
	publicKey1 := ImportECDSAPublic(publicKey)
	cryptText = xerror.PanicErr(Encrypt(rand.Reader, publicKey1, plainText, nil, nil)).([]byte)

	return
}

// The private key and plaintext are passed in for decryption
func EccDecrypt(cryptText, key []byte) (msg []byte, err error) {
	defer xerror.RespErr(&err)

	block, _ := pem.Decode(key)
	tempPrivateKey := xerror.PanicErr(x509.ParseECPrivateKey(block.Bytes)).(*ecdsa.PrivateKey)
	privateKey := ImportECDSA(tempPrivateKey)
	msg = xerror.PanicErr(privateKey.Decrypt(cryptText, nil, nil)).([]byte)

	return
}
