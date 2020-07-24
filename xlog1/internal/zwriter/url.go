package zwriter

import (
	"bytes"
	"net/http"
)

func NewUrlWriter(url string) *urlWriter {
	return &urlWriter{url: url}
}

type urlWriter struct {
	url string
}

func (t *urlWriter) Write(p []byte) (n int, err error) {
	res, err := http.Post(t.url, "application/json", bytes.NewBuffer(p))
	defer res.Body.Close()
	return 0, err
}
