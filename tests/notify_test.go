package tests

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gen2brain/beeep"
	"github.com/gen2brain/dlgs"
	"github.com/pubgo/g/xerror"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNotify(t *testing.T) {
	watcher := xerror.PanicErr(fsnotify.NewWatcher()).(*fsnotify.Watcher)
	defer watcher.Close()

	xerror.Panic(filepath.Walk(os.ExpandEnv("$GOPATH/src/mydocs/typora"),
		func(path string, info os.FileInfo, err1 error) (err error) {
			defer xerror.RespErr(&err)

			xerror.Panic(err1)

			if info.IsDir() {
				return
			}

			xerror.Panic(watcher.Add(path))

			return nil
		}))

	for {
		select {
		case event, ok := <-watcher.Events:
			if ok {
				fmt.Println(event.String())
			}

			//if ok && (event.Op == fsnotify.Write || event.Op == fsnotify.Rename) {
			//	_file := event.Name
			//	fmt.Println(_file)
			//xerror.Panic(watcher.Add(_file))
			//}
		case err, ok := <-watcher.Errors:
			if ok {
				xerror.Panic(err)
			}
		}
	}
}

func TestBeeep(t *testing.T) {
	xerror.Panic(beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration))

	//xerror.Panic(beeep.Notify("Title", "Message body", os.ExpandEnv("$PWD/assets/information.png")))

	xerror.Panic(beeep.Alert("Title", "Message body", os.ExpandEnv("$PWD/assets/warning.png")))
}

func TestDlgs(t *testing.T) {
	dlgs.Date("ss", "dd", time.Now())

	dlgs.File("dd", "", false)

	dlgs.Entry("dd", "dd", "dd")

	dlgs.Error("ddd", "ddd")

	dlgs.Color("dd", "")

	item, _, err := dlgs.List("List", "Select item from list:", []string{"Bug", "New Feature", "Improvement"})
	if err != nil {
		panic(err)
	}
	fmt.Println(item)

	passwd, _, err := dlgs.Password("Password", "Enter your API key:")
	if err != nil {
		panic(err)
	}
	fmt.Println(passwd)

	yes, err := dlgs.Question("Question", "Are you sure you want to format this media?", true)
	if err != nil {
		panic(err)
	}
	fmt.Println(yes)

	fmt.Println(dlgs.Info("titile", "content"))
}
