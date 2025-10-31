// Package ioext has misc extensions related to IO.
package ioext

import (
	"fmt"
	"io"
)

type countingWriter struct {
	w     io.Writer
	count int64
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	cw.count += int64(n)
	if err != nil {
		return n, fmt.Errorf("counting Write failed: %w", err)
	}
	return n, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &countingWriter{w: w}
	return cw, &cw.count
}
