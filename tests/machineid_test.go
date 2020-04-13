package tests

import (
	"fmt"
	"github.com/denisbrodbeck/machineid"
	"github.com/pubgo/g/pkg/encoding/baseutil"
	"github.com/pubgo/g/xerror"
	"github.com/rs/xid"
	"log"
	"testing"
)

func TestMachineid(t *testing.T) {
	id, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)

	id, err = machineid.ProtectedID("myAppName")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}

func TestId(t *testing.T) {
	guid := xid.New()

	fmt.Println(
		guid.String(),
		string(xerror.PanicBytes(guid.MarshalText())),
		string(guid.Machine()),
		guid.Pid(),
		guid.Time(),
		guid.Counter(),
	)

	_t := xid.NilID()
	xerror.Panic(_t.UnmarshalText([]byte("boevrdth5s5m00ppm0ug")))
	fmt.Println(_t.Machine())

	_id := xerror.PanicErr(xid.FromString("boevrdth5s5m00ppm0ug")).(xid.ID)
	fmt.Println(baseutil.Base62.EncodeToString(_id.Bytes()))
}
