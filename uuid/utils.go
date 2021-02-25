package uuid

import (
	"github.com/pubgo/x/net2"
	"github.com/pubgo/x/strutil"
	"github.com/sony/sonyflake"
	"github.com/speps/go-hashids"
)

var sf *sonyflake.Sonyflake
var upperMachineID uint16

func init() {
	var st sonyflake.Settings
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		sf = sonyflake.NewSonyflake(sonyflake.Settings{
			MachineID: net2.Lower16BitIP,
		})
		upperMachineID, _ = net2.Upper16BitIP()
	}
}

func GetIntId() uint64 {
	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}
	return id
}

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

const Alphabet36 = "abcdefghijklmnopqrstuvwxyz1234567890"

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
