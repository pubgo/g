package zwriter

import (
	"io"
)

func NewZWriter() *ZWriter {
	z := &ZWriter{
	}
	return z
}

// 可以是console,url,nsq,redis,...
type ZWriter struct {
	writer []io.Writer
}

func (t *ZWriter) Append(writer io.Writer) {
	t.writer = append(t.writer, writer)
}

func (t *ZWriter) Write(p []byte) (n int, err error) {
	return t.write(p)
}

func (t *ZWriter) write(p []byte) (n int, err error) {
	for _, w := range t.writer {
		if n, err := w.Write(p); err != nil {
			return n, err
		}
	}

	return 0, nil
}
