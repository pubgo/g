package ethutil

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"strings"
)

func Sign(data []byte, privateKey string) (string, error) {
	locPrivateKey, err := crypto.HexToECDSA(privateKey)

	sigStr := crypto.Keccak256(data)
	sigBytes, err := crypto.Sign(sigStr, locPrivateKey)

	return hex.EncodeToString(sigBytes), err
}

func Verify(data []byte, signature, public string) bool {
	msgBytes := crypto.Keccak256(data)

	sigBytes, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	publicKey, err := secp256k1.RecoverPubkey(msgBytes, sigBytes)
	if err != nil {
		return false
	}
	locPublicKey, err := crypto.UnmarshalPubkey(publicKey)
	if err != nil {
		return false
	}
	locAddress := crypto.PubkeyToAddress(*locPublicKey)

	return strings.ToLower(locAddress.Hex()) == strings.ToLower(public)
}

func SignByEth(data []byte, privateKey *ecdsa.PrivateKey) (string, error) {
	sigStr := crypto.Keccak256(data)
	sigBytes, err := crypto.Sign(sigStr, privateKey)

	return hex.EncodeToString(sigBytes), err
}
