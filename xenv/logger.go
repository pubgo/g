package xenv

import (
	"fmt"
	"log"
	"os"
)

var logger = log.New(os.Stderr, "xenv: ", log.LstdFlags)

func fatal(v error) {
	if v == nil {
		return
	}

	s := fmt.Sprintln(v)
	logger.Output(2, fmt.Sprintln(v))
	panic(s)
}
