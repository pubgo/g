package rand2

import (
	"encoding/hex"
)

func String(n int) string {
	return hex.EncodeToString(Bytes(n))
}
