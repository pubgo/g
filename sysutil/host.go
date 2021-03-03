package sysutil

import (
	"github.com/pubgo/x/strutil"
	
	"os"
)

var hostname string

func init() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		hostname = strutil.RandId()
	}
}

func Hostname() string {
	return hostname
}
