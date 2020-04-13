// +build !jsoniter

package json

import "encoding/json"

var (
	MarshalIndent = json.MarshalIndent
	NewDecoder    = json.NewDecoder
	NewEncoder    = json.NewEncoder
	Marshal       = json.Marshal
	Unmarshal     = json.Unmarshal
)
