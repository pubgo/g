package shutil

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pubgo/x/strutil"
	"github.com/pubgo/xerror"
)

func Exec(shell ...string) (string, error) {
	var out = strutil.GetBuilder()
	defer out.Reset()

	cmd := Bash(shell...)
	cmd.Stdout = out

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return out.String(), nil
}

func Run(cmd *exec.Cmd) (string, error) {
	var out = strutil.GetBuilder()
	defer out.Reset()

	cmd.Stdout = out

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return out.String(), nil
}

func Cmd(args ...string) *exec.Cmd {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd
}

func Bash(args ...string) *exec.Cmd {
	cmd := exec.Command("/bin/bash", "-c", strings.Join(args, " "))
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd
}

func GoMod() (string, error) {
	return Exec("go mod graph")
}

func GoList() (string, error) {
	return Exec("go list ./...")
}

func GraphViz(in, out string) (err error) {
	defer xerror.RespErr(&err)

	ret, err := Exec("dot", "-Tsvg", in)
	xerror.PanicF(err, "in:%s, out:%s", in, out)

	return ioutil.WriteFile(out, []byte(ret), 0600)
}

// NtpTimeDiff returns the time difference between the localhost and the given NTP server.
func NtpTimeDiff(server string) (time.Duration, error) {
	output, err := Run(Cmd("/usr/sbin/ntpdate", "-q", server))
	if err != nil {
		return time.Duration(0), err
	}

	re, _ := regexp.Compile("offset (.*) sec")
	matched := re.FindSubmatch([]byte(output))
	if len(matched) != 2 {
		return time.Duration(0), errors.New("invalid ntpdate output")
	}

	f, err := strconv.ParseFloat(string(matched[1]), 64)
	if err != nil {
		return time.Duration(0), err
	}

	return time.Duration(f*1000) * time.Millisecond, nil
}
