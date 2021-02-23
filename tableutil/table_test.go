package tableutil_test

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"github.com/pubgo/x/tableutil"
	"testing"
)

func TestName(t *testing.T) {

	tbl := tableutil.New("ID", "Name", "Score", "Added").
		WithPadding(1).
		WithWidthFunc(runewidth.StringWidth).
		WithFirstColumnFormatter(func(s string, i ...interface{}) string {
		return fmt.Sprintf(s,i...)
	})
	tbl.AddRow("ss", "ssssssssss", "s", "")
	tbl.AddRow("ss", "ssssssssss", "s", "")
	tbl.AddRow("ss", "ssssssssss", "s", "")
	tbl.AddRow("ssssssssss", "ssss", "s", "")
	tbl.Print()
}
