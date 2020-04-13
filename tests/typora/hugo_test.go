package tests

import (
	"github.com/pubgo/g/models"
	"github.com/pubgo/g/pkg/fileutil"
	"github.com/pubgo/g/xerror"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// 通过dbmeta里面的元数据，生成hugo 文章目录，然后发布
func TestHugoGen(t *testing.T) {
	_root := "/Users/barry/gopath/src/mydocs"
	_meta := filepath.Join(_root, "dbmeta")
	_hugo := filepath.Join(_root, "content", "post")

	xerror.Panic(filepath.Walk(filepath.Join(_meta, "notes", "知识库"), func(path string, info os.FileInfo, err1 error) (err error) {
		defer xerror.RespErr(&err)

		xerror.Panic(err1)
		if info.IsDir() || info.Name() == ".DS_Store" {
			return
		}

		_m := &models.Matter{}
		xerror.Panic(yaml.Unmarshal([]byte(xerror.PanicStr(fileutil.ReadFile(path))), _m))
		_path := filepath.Join(_root, _m.Path)
		hm := &models.HugoMatter{}
		hm.ID = _m.ID
		hm.Title = _m.Title
		hm.Tags = _m.Tags
		hm.Categories = _m.Categories
		hm.Date = _m.Created
		hm.Lastmod = _m.Modified
		hm.Toc = true
		hm.Weight = _m.Weight
		hm.KeyWords = _m.KeyWords

		return ioutil.WriteFile(filepath.Join(_hugo, _m.ID)+".md", []byte(hm.String()+xerror.PanicStr(fileutil.ReadFile(_path))), 0644)
	}))
}
