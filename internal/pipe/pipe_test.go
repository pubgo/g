package pipe

import (
	"encoding/hex"
	"fmt"
	"github.com/pubgo/x/utilx"
	"math/big"
	"reflect"
	"strings"
	"testing"
)

type t1 struct {
	A string
	b int
}

func TestP(t *testing.T) {
	SliceOf([]int{1, 2, 3}, []int{1, 2, 3}).P()
	SliceOf(t1{A: "dd", b: 1}, &t1{A: "sss", b: 2}).P()
}

func TestFilter(t *testing.T) {
	t.Run("test filter", func(t *testing.T) {
		SliceOf(t1{A: "dd", b: 1}, &t1{A: "sss", b: 2}).Filter(func(i int, v interface{}) bool {
			return !utilx.IsZero(v)
		}).P()
	})

	t.Run("test filter type", func(t *testing.T) {
		SliceOf(&t1{A: "dd", b: 1}, &t1{A: "sss", b: 2}).Filter(func(v *t1) bool {
			return v.b > 1
		}).P()
	})

	Filter([]string{"1", "2"}, func(v string) bool {
		return v == "1"
	}).P()
}

func TestMap(t *testing.T) {
	SliceOf(t1{A: "dd", b: 1}, t1{A: "sss", b: 2}).Map(func(i int, v t1) t1 {
		v.b = 100000
		return v
	}).P("test map")

	SliceOf(t1{A: "dd", b: 1}, t1{A: "sss", b: 2}).Map(func(i int, v t1) t1 {
		v.b = 100000
		return v
	}).Each(func(i int, a ...t1) {
		fmt.Println(a)
	})

	SliceOf(t1{A: "dd", b: 1}, t1{A: "sss", b: 2}).Map(func(i int, v t1) t1 {
		v.b = 100000
		return v
	}).P()

	SliceOf(nil, &t1{}).Map(func(i int, v *t1) *t1 {
		if v == nil {
			return nil
		}

		fmt.Println(v.b)

		v.b = 100000

		return v
	}).Map(func(v *t1) *t1 {
		if v == nil {
			return nil
		}

		fmt.Println("map2", v.b)
		v.b = 222000000
		return v
	}).Each(func(v *t1) {
		fmt.Println(v)
	})
}

func TestArray(t *testing.T) {
	var ddd []int
	ddd = append(ddd, 1, 2, 34)
	ArrayOf(ddd).Each(func(i, n int) {
		fmt.Println(i, n)
	})
}

func TestReduce(t *testing.T) {

	SliceOf(t1{A: "dd", b: 1}, t1{A: "sss", b: 2}).Map(func(v t1) t1 {
		v.b = 100000
		return v
	}).Reduce(func(s t1, v t1) t1 {
		return t1{b: s.b + v.b, A: s.A + v.A}
	}).Each(func(a interface{}) {
		fmt.Println(a)
	})

	fmt.Println(SliceOf(t1{A: "dd", b: 1}, t1{A: "sss", b: 2}, t1{A: "sss", b: 2}).Map(func(i int, v t1) t1 {
		v.b = 100000
		return v
	}).Reduce(func(s t1, v t1) t1 {
		return t1{b: s.b + v.b, A: s.A + v.A}
	}).ToJSON())
}

func TestEach(t *testing.T) {
	SliceOf(1, 2, 3, 4).Each(func(a ...interface{}) {
		fmt.Println(a)
	})

	SliceOf(1, 2, 3, 4).Each(func(i int, a interface{}) {
		fmt.Println(i, a)
	})

	SliceOf(1, 2, 3, 4).Each(func(a interface{}) {
		fmt.Println(a)
	})
}

func TestPipe(t *testing.T) {
	SliceOf(1, "dd").Pipe(func(a int, b string) (int, string) {
		fmt.Println("callback success ok", a, b)
		return a, b
	}).Pipe(func(a int, b string) {
		fmt.Println("callback ", a, b)
	})

	SliceOf(1, 2, 3, 4, nil).Pipe(func(a, b, c, d int, e error) {
		fmt.Println(a, b, c, d, e)
	}).P("test pipe")
}

func TestError(t *testing.T) {
	//SliceOf()(1, 2, 3, errors.New("sss")).MustNotError()
	//SliceOf()(1, 2, 3, nil).MustNotNil()
}

func TestToData(t *testing.T) {
	a := "0 */2 * * * *"
	fmt.Println(SliceOf(strings.Split(a, "*")[1]).ToString())
	fmt.Println(ArrayOf(strings.Split(a, "*")).ToString())
	fmt.Println(ArrayOf(strings.Split(a, "*")).ToString())
	SliceOf(1, 2, 3, "", nil, &a).P()
	SliceOf(1, 2, 3, "", nil, &a).P()
}

func TestSetInterface(t *testing.T) {
	_fn := func(in interface{}, a interface{}) {
		fmt.Println(in, a)
		reflect.ValueOf(in).Elem().Set(reflect.ValueOf(a))
	}
	a := 1
	b := 2
	_fn(&a, b)
	fmt.Println(a, b)
}

func TestExpr(t *testing.T) {
	SliceOf(1, 2, 3, 4, nil).Pipe(func(a, b, c, d int, e error) {
		fmt.Println(a, b, c, d, e)
	}).P("test pipe")

	//fmt.Println(SliceOf()(1, 2, 3, 4, nil).FilterExp(`t != nil`).ToJson())
	//fmt.Println(SliceOf()(&t1{A: "1", b: 2}, &t1{A: "1", b: 3}).FilterExp(`t.A == "1"`).ToJson())
	//fmt.Println(SliceOf()(&t1{A: "1", b: 2}, &t1{A: "1", b: 3}).MapExp(`t.A == "1"`).ToJson())

	_a := SliceOf(nil, &t1{A: "1", b: 2}, &t1{A: "1", b: 3}).ToData().([]*t1)
	fmt.Println(_a)
	fmt.Println(_a[1].A)
}

type M struct {
	A  string `json:"a"`
	A1 string `json:"a1"`
	A2 string `json:"a2"`
}

func (t *M) Name() string {
	return "m"
}

func TestSortBy(t *testing.T) {
	if a, ok := SortBy([]string{"11", "2", "3"}, func(a, b string) bool {
		return strings.Compare(a, b) > 0
	}).ToData().([]string); ok {
		fmt.Println(a)
	}

	fmt.Println("nil test", SliceOf(nil, &t1{A: "1", b: 2}, &t1{A: "1", b: 3}).
		SortBy(func(a, b *t1) bool {
			if a == nil || b == nil {
				return true
			}

			return a.b > b.b
		}).ToData())

	fmt.Println(SliceOf(1, 11, 2).SortBy(func(a, b int) bool {
		return a > b
	}).ToJSON())

	SliceOf(1, 11, 2).SortBy(func(a, b int) bool {
		return a > b
	}).ToData(func(a []int) {
		fmt.Println(a)
	})
}

func TestGroupBy(t *testing.T) {
	SliceOf(nil, &t1{A: "1", b: 2}, map[string]interface{}{"A": "2"}, t1{A: "1", b: 2}, &t1{A: "2", b: 3}, &t1{A: "2", b: 3})
}

func TestName123(t *testing.T) {
	a := "1a38753878917df87acfac"
	b, err := hex.DecodeString(a)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(big.NewInt(0).SetBytes(b).Int64())
}
