package tests

import (
	"fmt"
	"github.com/pubgo/g/xerror"
	"regexp"
	"runtime/debug"
	"testing"
)

func TestEnv(t *testing.T) {
	var _envRegexp = xerror.PanicErr(regexp.Compile(`\${(.+)}`)).(*regexp.Regexp)

	fmt.Println(_envRegexp.MatchString(`${app_secret}`))
	fmt.Println(_envRegexp.FindStringSubmatch(`${app_secret}`))
	fmt.Println(_envRegexp.MatchString(`${hh||hello}`))
	fmt.Println(_envRegexp.FindStringSubmatch(`${hh||hello}`))

	var _safeEnvRegexp = xerror.PanicErr(regexp.Compile(`!{(.+)}`)).(*regexp.Regexp)
	fmt.Println(_safeEnvRegexp.MatchString(`!{app_secret}`))
	fmt.Println(_safeEnvRegexp.FindStringSubmatch(`!{app_secret}`))
	fmt.Println(_safeEnvRegexp.MatchString(`!{hh||hello}`))
	fmt.Println(_safeEnvRegexp.FindStringSubmatch(`!{hh||hello}`))
	debug.PrintStack()
}
