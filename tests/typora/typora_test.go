package tests

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/pubgo/g/pkg/encoding/baseutil"
	"github.com/pubgo/g/pkg/fileutil"
	"github.com/pubgo/g/pkg/strutil"
	"github.com/pubgo/g/pkg/timeutil"
	"github.com/pubgo/g/xerror"
	"github.com/rs/xid"
	"github.com/yanyiwu/gojieba"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

var trim = strings.TrimSpace
var titleReg = regexp.MustCompile(`.*#(.+)`)

func TestTypora(t *testing.T) {
	defer xerror.Assert()

	xerror.Panic(filepath.Walk(os.ExpandEnv("$GOPATH/src/mydocs/typora"),
		func(path string, info os.FileInfo, err1 error) (err error) {
			defer xerror.RespErr(&err)

			xerror.Panic(err1)

			if info.IsDir() {
				return
			}

			if info.Size() == 0 {
				xerror.Panic(fileutil.DeleteFile(path))
				return
			}

			if trim(info.Name())[0] == '!' {
				fmt.Println(path)
				for _, _dt := range fileutil.ReadLine(path) {
					_dt = trim(_dt)
					fmt.Println(_dt)
					if strings.HasPrefix(_dt, "#") {
						//_title := titleReg.FindStringSubmatch(_dt)[1]
						fmt.Println(filepath.Dir(path), _dt)
						//xerror.Panic(os.Rename(path, filepath.Join(filepath.Dir(path), _title+".md")))
						//os.Rename()
						break
					}
				}

			}

			return nil
		}))
}

type Matter struct {
	Tags       []string `yaml:",flow"`
	KeyWords   []string `yaml:",flow"`
	Title      string   `yaml:",omitempty"`
	Created    string   `yaml:",omitempty"`
	Modified   string   `yaml:",omitempty"`
	IsDraft    bool
	Published  string   `yaml:",omitempty"`
	Categories []string `yaml:",flow,omitempty"`
	WordCount  int      `yaml:",omitempty"`
	Abstract   string   `yaml:",omitempty"`
	Format     string   `yaml:",omitempty"`
	Kind       string   `yaml:",omitempty"`
	Weight     int      `yaml:",omitempty"`
	ID         string   `yaml:",omitempty"`
	Comments   bool     `yaml:",omitempty"`
	MinHash    string   `yaml:",omitempty"`
	Namespaces string   `yaml:",omitempty"`
	Path       string   `yaml:",omitempty"`
	Platform   map[string]struct {
		ID        string `yaml:",omitempty"`
		Published string `yaml:",omitempty"`
		Updated   string `yaml:",omitempty"`
		URL       string `yaml:",omitempty"`
	} `yaml:",omitempty"`
}

// 元数据处理
func TestFormater(t *testing.T) {
	x := gojieba.NewJieba()
	defer x.Free()

	_root := os.ExpandEnv("/Users/barry/gopath/src/mydocs")
	_meta := filepath.Join(_root, "dbmeta")
	xerror.Panic(filepath.Walk(filepath.Join(_root, "notes"),
		func(path string, info os.FileInfo, err1 error) (err error) {
			defer xerror.RespErr(&err)

			xerror.Panic(err1)

			if info.IsDir() || info.Name() == ".DS_Store" {
				return
			}

			if info.Size() == 0 {
				xerror.Panic(fileutil.DeleteFile(path))
				return
			}

			_fd := strings.TrimSpace(xerror.PanicStr(fileutil.ReadFile(path)))
			_dt := strings.Split(_fd, "\n")
			//if !strings.HasPrefix(_dt[0], "---") {
			//	return
			//}

			var title = strings.ReplaceAll(info.Name(), filepath.Ext(info.Name()), "")

			for _, _d := range _dt {
				_d = trim(_d)
				if strings.HasPrefix(_d, "#") {
					if _d != "" {
						title = _d
					}
					break
				}
			}

			title = strings.Trim(title, "'# ")

			_s1 := strings.Split(strings.ReplaceAll(path, _root, ""), "/")
			_s1 = _s1[:len(_s1)-1]

			_namespaces := strings.Join(strutil.Filter(_s1, ""), "/")

			_matter := string(xerror.PanicBytes(yaml.Marshal(
				&Matter{
					Tags:       strutil.Filter(_s1, ""),
					Title:      title,
					Created:    info.ModTime().Format(timeutil.Format),
					Modified:   info.ModTime().Format(timeutil.Format),
					Categories: strutil.Filter(_s1, ""),
					WordCount:  len([]rune(_fd)),
					Namespaces: _namespaces,
					Path:       strings.ReplaceAll(path, _root, ""),
					ID:         baseutil.Base62.EncodeToString(xid.New().Bytes()),
					KeyWords:   x.Extract(_fd, 5),
					Weight:     1,
				},
			)))

			_p := filepath.Join(_meta, strings.ReplaceAll(path, _root, ""))
			xerror.Panic(fileutil.IsNotExistMkDir(filepath.Dir(_p)))
			return ioutil.WriteFile(_p+".yaml", []byte(_matter), 0644)
		}))
}

func TestMeta(t *testing.T) {
	_root := "/Users/barry/gopath/src/mydocs/mweb/metadata"
	//var media = make(map[string]interface{})
	xerror.Panic(filepath.Walk(_root,
		func(path string, info os.FileInfo, err1 error) (err error) {
			defer xerror.RespErr(&err)

			xerror.Panic(err1)
			if info.IsDir() || info.Name() == ".DS_Store" {
				return
			}

			if info.Size() == 0 {
				xerror.Panic(fileutil.DeleteFile(path))
				return
			}

			var dt = make(map[string]interface{})
			if json.Unmarshal([]byte(xerror.PanicStr(fileutil.ReadFile(path))), &dt) != nil {
				fmt.Println(xerror.PanicStr(fileutil.ReadFile(path)))
				return nil
			}

			//if _dt, ok := dt["medias"].(map[string]interface{}); ok {
			//	for k, v := range _dt {
			//		media[k] = v
			//	}
			//}

			return nil
		},
	))

	//fmt.Println(string(xerror.PanicBytes(json.Marshal(media))))
}

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

func TestCsv(t *testing.T) {
	bs := xerror.PanicBytes(ioutil.ReadFile("/Users/barry/gopath/src/mydocs/csv/article.csv"))

	x := gojieba.NewJieba()
	defer x.Free()

	_root := os.ExpandEnv("/Users/barry/gopath/src/mydocs")
	_meta := filepath.Join(_root, "dbmeta")

	_t := &Table{}
	xerror.Panic(_t.LoadCSV(bytes.NewReader(bs)))

	var ss = make(map[string][]map[string]interface{})
	xerror.Panic(_t.Each(func(i map[string]interface{}) error {
		ss[i["aid"].(string)] = append(ss[i["aid"].(string)], i)
		return nil
	}))

	for k, v := range ss {
		_path := "/Users/barry/gopath/src/mydocs/notes/mweb/" + k + ".md"
		if !fileutil.IsExist(_path) {
			continue
		}

		_fd := xerror.PanicStr(fileutil.ReadFile(_path))
		_dt := strings.Split(_fd, "\n")
		var title = k

		for _, _d := range _dt {
			_d = trim(_d)
			if strings.HasPrefix(_d, "#") {
				if _d != "" {
					title = _d
				}
				break
			}
		}

		title = strings.Trim(title, "'# ")
		_m := &Matter{
			ID:         baseutil.Base62.EncodeToString(xid.New().Bytes()),
			Title:     title,
			WordCount: len([]rune(_fd)),
			Path:      strings.ReplaceAll(_path, _root, ""),
			KeyWords:  x.Extract(_fd, 5),
		}
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

		_p := filepath.Join(_meta, strings.ReplaceAll(_path, _root, ""))
		xerror.Panic(fileutil.IsNotExistMkDir(filepath.Dir(_p)))
		xerror.Panic(ioutil.WriteFile(_p+".yaml", []byte(string(xerror.PanicBytes(yaml.Marshal(_m)))), 0644))
	}
}

