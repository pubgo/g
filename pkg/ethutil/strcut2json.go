package ethutil

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/shopspring/decimal"
)

const (
	loc_fieldnamej_signature = "signature"
	loc_fieldnamej_content   = "content"
)

func init() {
	decimal.MarshalJSONWithoutQuotes = true
}

func ToSignatureByExclude(obj interface{}, excludes ...string) (string, error) {
	beforeValue, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	midResultMap := make(map[string]interface{}, 0)

	jen := json.NewDecoder(bytes.NewBuffer(beforeValue))
	jen.UseNumber()
	err = jen.Decode(&midResultMap)
	if err != nil {
		return "", err
	}

	delete(midResultMap, loc_fieldnamej_signature)

	for _, item := range excludes {
		delete(midResultMap, item)
	}

	eValue, err := JsonMarshal(midResultMap)
	if err != nil {
		return "", err
	}

	return string(eValue), nil
}

func RecoverPublickeyBySignature(msg, sig string) (string, error) {
	msgBytes := crypto.Keccak256([]byte(msg))

	sigBytes, err := hex.DecodeString(sig)
	if err != nil {
		return "", err
	}

	publicKey, err := secp256k1.RecoverPubkey(msgBytes, sigBytes)
	if err != nil {
		return "", err
	}

	srcPublicKey, _ := crypto.UnmarshalPubkey(publicKey)
	srcAddress := crypto.PubkeyToAddress(*srcPublicKey)

	return srcAddress.Hex(), nil
}

// The default behavior is to escape &, <, and > to \u0026, \u003c, and \u003e
// to avoid certain safety problems that can arise when embedding JSON in HTML.
func JsonMarshal(value map[string]interface{}) (string, error) {
	buffer := &bytes.Buffer{}

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(value)
	if err != nil {
		return "", err
	}

	result := strings.Trim(buffer.String(), "\n")

	return result, nil
}
