package iox

import (
	"context"
	"io"
	"sync"
)

func Copy(in io.Reader, out io.Writer) (err error) {
	var (
		count int
		buf   = make([]byte, 8192)
	)

	for {
		count, err = in.Read(buf)
		if err != nil {
			if err == io.EOF && count > 0 {
				_, _ = out.Write(buf[:count])
			}
			break
		}

		if count > 0 {
			_, err = out.Write(buf[:count])
		}
	}
	return
}

func ReadAndWrite(ctx context.Context, r io.Reader, w io.Writer, wg *sync.WaitGroup) <-chan error {
	c := make(chan error)
	go func() {
		if wg != nil {
			defer wg.Done()
		}

		buff := make([]byte, 1024)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				count, err := r.Read(buff)
				if err != nil {
					return
				}

				if count > 0 {
					_, err := w.Write(buff[:count])
					if err != nil {
						return
					}
				}
			}
		}
	}()
	return c
}
