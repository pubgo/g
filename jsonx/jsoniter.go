package jsonx

import "github.com/json-iterator/go"

var (
	json2         = jsoniter.ConfigCompatibleWithStandardLibrary
	MarshalIndent = json2.MarshalIndent
	NewDecoder    = json2.NewDecoder
	NewEncoder    = json2.NewEncoder
	Marshal       = json2.Marshal
	Unmarshal     = json2.Unmarshal
)
