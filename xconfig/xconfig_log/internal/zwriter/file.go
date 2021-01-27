package zwriter

import (
	"github.com/pubgo/x/xerror"
	"os"
)

func NewFileWriter(path string) *fileWriter {
	return &fileWriter{f: xerror.PanicErr(os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)).(*os.File)}
}

type fileWriter struct {
	f *os.File
}

func (t *fileWriter) Write(p []byte) (n int, err error) {
	return t.f.Write(p)
}

func (t *fileWriter) Close() error {
	return t.f.Close()
}

// Rotate renames old log file, creates new one, switches log and closes the old file.
func (t *fileWriter) Rotate() error {
	// rename dest file if it already exists.
	//if _, err := os.Stat(l.name); err == nil {
	//	name := l.name + "." + time.Now().Format(time.RFC3339)
	//	if err = os.Rename(l.name, name); err != nil {
	//		return err
	//	}
	//}
	//// create new file.
	//file, err := os.Create(l.name)
	//if err != nil {
	//	return err
	//}
	//// switch dest file safely.
	//l.mu.Lock()
	//file, l.file = l.file, file
	//l.mu.Unlock()
	//// close old file if open.
	//if file != nil {
	//	if err := file.Close(); err != nil {
	//		return err
	//	}
	//}
	return nil
}
