package hashutil

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

// MD5 md5 encryption
func MD5(data string) string {
	m := md5.New()
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil))
}

// MD5Verify md5 verify
func MD5Verify(data string, v string) bool {
	m := md5.New()
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil)) == v
}

// Keccak256 calculates and returns the Keccak256 hash of the input data.
func Keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

// Keccak512 calculates and returns the Keccak512 hash of the input data.
func Keccak512(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak512()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

func Sha256(data ...[]byte) []byte {
	_hash := sha256.New()
	for _, b := range data {
		_hash.Write(b)
	}
	return _hash.Sum(nil)
}

func Sha512(data ...[]byte) []byte {
	_hash := sha512.New()
	for _, b := range data {
		_hash.Write(b)
	}
	return _hash.Sum(nil)
}

func Sha1(data ...[]byte) []byte {
	_hash := sha1.New()
	for _, b := range data {
		_hash.Write(b)
	}
	return _hash.Sum(nil)
}
