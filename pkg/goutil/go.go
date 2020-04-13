package goutil

import (
	"github.com/pubgo/g/pkg/fileutil"
	"github.com/pubgo/g/pkg/osutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

const (
	pathSeparator     = string(os.PathSeparator)     // OS-specific path separator
	pathListSeparator = string(os.PathListSeparator) // OS-specific path list separator
)

// GetAPIPath gets the Go source code path $GOROOT/src.
func GetAPIPath() string {
	return filepath.FromSlash(path.Clean(runtime.GOROOT() + "/src"))
}

// IsAPI determines whether the specified path belongs to Go API.
func IsAPI(path string) bool {
	apiPath := GetAPIPath()

	return strings.HasPrefix(filepath.FromSlash(path), apiPath)
}

// GetGoFormats gets Go format tools. It may return ["gofmt", "goimports"].
func GetGoFormats() []string {
	ret := []string{"gofmt"}

	p := GetExecutableInGOBIN("goimports")
	if fileutil.IsExist(p) {
		ret = append(ret, "goimports")
	}

	sort.Strings(ret)

	return ret
}

// GetExecutableInGOBIN gets executable file under GOBIN path.
//
// The specified executable should not with extension, this function will append .exe if on Windows.
func GetExecutableInGOBIN(executable string) string {
	if osutil.IsWindows() {
		executable += ".exe"
	}

	gopaths := filepath.SplitList(os.Getenv("GOPATH"))

	for _, gopath := range gopaths {
		// $GOPATH/bin/$GOOS_$GOARCH/executable
		ret := gopath + pathSeparator + "bin" + pathSeparator +
			os.Getenv("GOOS") + "_" + os.Getenv("GOARCH") + pathSeparator + executable
		if fileutil.IsExist(ret) {
			return ret
		}

		// $GOPATH/bin/{runtime.GOOS}_{runtime.GOARCH}/executable
		ret = gopath + pathSeparator + "bin" + pathSeparator +
			runtime.GOOS + "_" + runtime.GOARCH + pathSeparator + executable
		if fileutil.IsExist(ret) {
			return ret
		}

		// $GOPATH/bin/executable
		ret = gopath + pathSeparator + "bin" + pathSeparator + executable
		if fileutil.IsExist(ret) {
			return ret
		}
	}

	// $GOBIN/executable
	gobin := os.Getenv("GOBIN")
	if "" != gobin {
		ret := gobin + pathSeparator + executable
		if fileutil.IsExist(ret) {
			return ret
		}
	}

	return "./" + executable
}
