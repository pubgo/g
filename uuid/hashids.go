package uuid

import (
	"github.com/pubgo/x/strutil"
	"github.com/speps/go-hashids"
)

// Hashids是一个小型的开源库，它从数字生成简短的、惟一的、非顺序的id。
// 它将像347这样的数字转换成像“yr8”这样的字符串，或者像[27,986]这样的数字数组转换成“3kTMd”。
// 您还可以解码这些id。
// 这对于将多个参数绑定到一个或简单地将它们用作简短的uid非常有用。

const Alphabet36 = "abcdefghijklmnopqrstuvwxyz1234567890"

// format likes: B6BZVN3mOPvx
func GetUuid(prefix string) string {
	id := GetIntId()
	hd := hashids.NewData()
	h, err := hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}

	i, err := h.Encode([]int{int(id)})
	if err != nil {
		panic(err)
	}

	return prefix + strutil.Reverse(i)
}

// format likes: 300m50zn91nwz5
func GetUuid36(prefix string) string {
	id := GetIntId()
	hd := hashids.NewData()
	hd.Alphabet = Alphabet36
	h, err := hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}
	i, err := h.Encode([]int{int(id)})
	if err != nil {
		panic(err)
	}

	return prefix + strutil.Reverse(i)
}
