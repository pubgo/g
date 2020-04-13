package tests

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pubgo/g/models"
	"github.com/pubgo/g/pkg/fileutil"
	"github.com/pubgo/g/xerror"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestMwebM(t *testing.T) {
	_root := "/Users/barry/gopath/src/mydocs/dbmeta/notes/mweb"
	_root1 := "/Users/barry/gopath/src/mydocs"

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

			_fd := xerror.PanicStr(fileutil.ReadFile(path))
			var sss = &models.Matter{}
			xerror.Panic(yaml.Unmarshal([]byte(_fd), sss))

			_fd1 := xerror.PanicStr(fileutil.ReadFile(filepath.Join(_root1, sss.Path)))

			_h := "---\n"
			_h += _fd
			_h += "---\n\n\n"
			_h += _fd1

			return ioutil.WriteFile(filepath.Join(_root1, sss.Path), []byte(_h), 0644)
		}))
}

// 监控修改变化，修改元数据
func TestNote(t *testing.T) {
	_notes := "/Users/barry/gopath/src/mydocs/dbmeta/notes"
	//_root1 := "/Users/barry/gopath/src/mydocs"
	_w := xerror.PanicErr(fsnotify.NewWatcher()).(*fsnotify.Watcher)
	defer _w.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-_w.Events:
				fmt.Println(event, ok)
				if !ok {
					return
				}

				log.Println("modified file:", event.Name)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}

				//_fd := xerror.PanicStr(fileutil.ReadFile(path))
				//var sss = &models.Matter{}
				//xerror.Panic(yaml.Unmarshal([]byte(_fd), sss))
				//
				//_fd1 := xerror.PanicStr(fileutil.ReadFile(filepath.Join(_root1, sss.Path)))
				//
				//_h := "---\n"
				//_h += _fd
				//_h += "---\n\n\n"
				//_h += _fd1
				//
				//return ioutil.WriteFile(filepath.Join(_root1, sss.Path), []byte(_h), 0644)
				//
				//log.Println("event:", event)

			case err, ok := <-_w.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	xerror.Panic(filepath.Walk(_notes,
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

			return _w.Add(path)
		}))
	<-done
}
