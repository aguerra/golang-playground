package ioext

import (
	"errors"
	"io"
	"testing"
)

type errWriter struct{}

func (w errWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("write failed")
}

func TestCountingWriter(t *testing.T) {
	inputs := []string{"", "one", "two", ""}
	cw, count := CountingWriter(io.Discard)
	var sum int64
	for _, input := range inputs {
		data := []byte(input)
		sum += int64(len(data))
		if _, err := cw.Write(data); sum != *count || err != nil {
			t.Errorf("Write(%q) = _ %v, count %v, want %v", input, err, *count, sum)
		}
	}
}

func TestCountingWriterErrWriter(t *testing.T) {
	w := errWriter{}
	cw, _ := CountingWriter(w)
	input := []byte("hello")
	if _, err := cw.Write(input); err == nil {
		t.Errorf("Write(%q) = _ %v, want error ", input, err)
	}
}
