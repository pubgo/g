package uuid

import (
	"github.com/pubgo/x/netutil"
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		sf = sonyflake.NewSonyflake(sonyflake.Settings{
			MachineID: netutil.Lower16BitIP,
		})
	}
}

func GetIntId() uint64 {
	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}
	return id
}
