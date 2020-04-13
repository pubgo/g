package shutil

import (
	"bytes"
	"os/exec"
	"sync"
)

var _b = &sync.Pool{
	New: func() interface{} {
		return bytes.NewBufferString("")
	},
}

func Execute(shell string) (string, error) {
	var out = _b.Get().(*bytes.Buffer)
	defer _b.Put(out)
	defer out.Reset()

	cmd := exec.Command("/bin/bash", "-c", shell)
	cmd.Stdout = out

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return out.String(), nil
}
