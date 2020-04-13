package csv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/pubgo/g/models"
	"github.com/pubgo/g/pkg/strutil"
	"github.com/pubgo/g/xerror"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"testing"
)

func ToMap(a1 []string, a2 []string) map[string]interface{} {
	var _m = make(map[string]interface{})
	for i := 0; i < len(a1); i++ {
		_m[a1[i]] = a2[i]
	}
	return _m
}

type Table struct {
	data []map[string]interface{}
}

func (t *Table) LoadCSV(rs io.ReadSeeker) error {
	cr := csv.NewReader(rs)
	//cr.ReuseRecord = true
	head, err := cr.Read()
	if err != nil {
		if err == io.EOF {
			rs.Seek(0, io.SeekStart)
			return nil
		}
		return err
	}

	for {
		rr, err := cr.Read()
		if err != nil {
			if err == io.EOF {
				rs.Seek(0, io.SeekStart)
				return nil
			}
			return err
		}

		t.data = append(t.data, ToMap(head, rr))
	}
}

func (t *Table) Each(fn func(map[string]interface{}) error) error {
	for _, d := range t.data {
		if err := fn(d); err != nil {
			return err
		}
	}
	return nil
}

type Matter = models.Matter

func TestCsv(t *testing.T) {
	bs := xerror.PanicBytes(ioutil.ReadFile("/Users/barry/gopath/src/mydocs/csv/article.csv"))

	_t := &Table{}
	xerror.Panic(_t.LoadCSV(bytes.NewReader(bs)))

	var ss = make(map[string][]map[string]interface{})
	xerror.Panic(_t.Each(func(i map[string]interface{}) error {
		ss[i["aid"].(string)] = append(ss[i["aid"].(string)], i)
		return nil
	}))

	for k, v := range ss {
		_m := &Matter{ID: k}
		for _, vv := range v {
			if !strutil.Contains(_m.Categories, vv["ns"].(string)) {
				_m.Categories = append(_m.Categories, vv["ns"].(string))
			}

			if !strutil.Contains(_m.Tags, vv["tag"].(string)) {
				_m.Tags = append(_m.Tags, vv["tag"].(string))
			}

			_m.Created = vv["dateAdd"].(string)
			_m.Modified = vv["dateModif"].(string)
			_m.Created = vv["dateArt"].(string)
			if vv["sort"].(string) != "" {
				_m.Weight = xerror.PanicErr(strutil.ToInt(vv["sort"].(string))).(int)
			}
			_m.IsDraft = vv["state"].(string) == "1"

		}
		fmt.Println(string(xerror.PanicBytes(yaml.Marshal(_m))))
	}
}
