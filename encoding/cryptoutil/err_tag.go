package cryptoutil

import (
	"errors"
)

// ErrTag err tags
var ErrTag = struct {
	ErrCipherKey           error
	ErrKeyLengthSixteen    error
	ErrKeyLengtheEight     error
	ErrKeyLengthTwentyFour error
	ErrPaddingSize         error
	ErrIvAes               error
	ErrIvDes               error
}{
	ErrCipherKey:           errors.New("The secret key is wrong and cannot be decrypted. Please check"),
	ErrKeyLengthSixteen:    errors.New("a sixteen or twenty-four or thirty-two length secret key is required"),
	ErrKeyLengtheEight:     errors.New("a eight-length secret key is required"),
	ErrKeyLengthTwentyFour: errors.New("a twenty-four-length secret key is required"),
	ErrPaddingSize:         errors.New("padding size error please check the secret key or iv"),
	ErrIvAes:               errors.New("a sixteen-length ivaes is required"),
	ErrIvDes:               errors.New("a eight-length ivdes key is required"),
}
