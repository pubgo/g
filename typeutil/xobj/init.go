// doc https://godoc.org/github.com/stretchr/objx
package xobj

import "github.com/stretchr/objx"

type (
	Map = objx.Map
	Value = objx.Value
	MSIConvertable = objx.MSIConvertable
)

const (
	PathSeparator      = objx.PathSeparator
	SignatureSeparator = objx.SignatureSeparator
)

var (
	Nil                  = objx.Nil
	HashWithKey          = objx.HashWithKey
	FromBase64           = objx.FromBase64
	FromJSON             = objx.FromJSON
	FromSignedBase64     = objx.FromSignedBase64
	FromURLQuery         = objx.FromURLQuery
	MSI                  = objx.MSI
	MustFromBase64       = objx.MustFromBase64
	MustFromJSON         = objx.MustFromJSON
	MustFromSignedBase64 = objx.MustFromSignedBase64
	MustFromURLQuery     = objx.MustFromURLQuery
	New                  = objx.New
)
