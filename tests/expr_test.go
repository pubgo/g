package tests

import (
	"fmt"
	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pubgo/g/xerror"
	"regexp"
	"strings"
	"testing"
	"time"
	"xorm.io/xorm"
)

type Request struct {
	Location  string
	Location1 int
	Date      time.Time
}

func (*Request) Join(a string) string {
	fmt.Println("join", a)
	return a
}

func (*Request) TB1(a string) string {
	fmt.Println("join", a)
	return a
}

func (*Request) TB2(a string) string {
	fmt.Println("join", a)
	return a
}

func (*Request) Before(a, b string) bool {
	fmt.Println(a, "||", b)
	return a > b
}

func (*Request) Before1(a, b int) bool {
	fmt.Println(a, b, "Before1")
	return a > b
}

func TestEnv123(t *testing.T) {
	var program *vm.Program
	var err error
	program, err = expr.Compile(`Join(TB1(a=1||b=2&&v>=2) as a,Tb2(v=2) as b,a.a=b.b)`,
		expr.Env(&Request{}),
	)
	xerror.Panic(err)
	output, err := expr.Run(program, &Request{Location: "wwwww", Location1: 12345})
	xerror.Panic(err)
	fmt.Println(output)
}

func TestSql(t *testing.T) {
	var _join = regexp.MustCompile(`(join\((.+)\))`)
	var _tbReg = regexp.MustCompile(`(\S.+)\((.*)\)`)
	var _tbAsReg = regexp.MustCompile(`(\S.+)\((.*)\)\sas\s(.+)`)
	a1 := func(a string) {
		if _join.MatchString(a) {
			fmt.Println("_join", _join.FindStringSubmatch(a))
			return
		}

		if _tbAsReg.MatchString(a) {
			fmt.Println("_tbAsReg", _tbAsReg.FindStringSubmatch(a))
			return
		}

		if _tbReg.MatchString(a) {
			fmt.Println("_tbReg", _tbReg.FindStringSubmatch(a))
			return
		}

		a = strings.ReplaceAll(a, "||", " or ")
		a = strings.ReplaceAll(a, "&&", " and ")
		fmt.Println(a)

	}

	a1(`join(TB1(a=1||b=2&&v>=2) as a,Tb2(v=2) as b,a.a=b.b)`)
	a1(`TB1(a=1||b=2&&v>=2)`)
	a1(`TB1(a=1||b=2&&v>=2) as a`)
	a1(`a=1||b=2&&v>=2`)
}

func TestXOrm(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	xerror.Panic(err)
	fmt.Println(engine)
}

/*
Join Union UnionAll

delete /res/:name Join(TB1(a=1||b=2&&v>=2) as a,Tb2(v=2) as b,a.a=b.b)
delete /res/:name TB1(a=1||b=2&&v>=2||v in [])

update /res/:name data Join(TB1(a=1||b=2&&v>=2) as a,Tb2(v=2) as b,a.a=b.b)

get /res/:name TB1(a=1||b=2&&v>=2||v nin [1,2,3,4])
get /res/:name Join(TB1(a=1||b=2&&v>=2) as a,Tb2(v=2) as b,a.a=b.b)
| group_by()
| sort(~a,~b)
| fields(a,b,c,d)
get /res/:name UnionAll(TB1(a=1||b=2&&v>=2) as a,Tb2(v=2) as b,a.a=b.b) | group_by(c1,c2,c3)
| sort(~a,~b)
| fields(count(a) as a1,b,c,d)
*/
