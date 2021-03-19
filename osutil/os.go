package osutil

import (
	"github.com/mitchellh/go-homedir"
)

// Home
// returns the home directory for the executing user.
func Home() (string, error) {
	return homedir.Dir()
}

func Expand(path string) (string, error) {
	return homedir.Expand(path)
}
