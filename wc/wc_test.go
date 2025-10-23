package wc

import (
	"errors"
	"strings"
	"testing"
	"testing/iotest"
)

func TestCountWords(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"", 0},
		{"one", 1},
		{"one two three", 3},
		{"one two three\tfour", 4},
		{"one two three\none two", 5},
	}
	for _, test := range tests {
		r := strings.NewReader(test.input)
		if got, err := CountWords(r); got != test.want || err != nil {
			t.Errorf("CountWords(%q) = %v %v, want %v", test.input, got, err, test.want)
		}
	}
}

func TestCountWordsErrReader(t *testing.T) {
	errReader := errors.New("reader error")
	r := iotest.ErrReader(errReader)
	if _, err := CountWords(r); !errors.Is(err, errReader) {
		t.Errorf("CountWords(ErrReader(%q)) = _ %v", errReader, err)
	}
}

func TestCountLines(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"", 0},
		{"\n", 1},
		{"one", 0},
		{"one\n", 1},
		{"\n\n", 2},
		{"one\ntwo", 1},
		{"one\ntwo\n", 2},
		{"\n\n\n", 3},
		{"one\ntwo\nthree", 2},
		{"one\ntwo\nthree\n", 3},
	}
	for _, test := range tests {
		r := strings.NewReader(test.input)
		if got, err := CountLines(r); got != test.want || err != nil {
			t.Errorf("CountLines(%q) = %v %v, want %v", test.input, got, err, test.want)
		}
	}
}

func TestCountLinesErrReader(t *testing.T) {
	errReader := errors.New("reader error")
	r := iotest.ErrReader(errReader)
	if _, err := CountLines(r); !errors.Is(err, errReader) {
		t.Errorf("CountLines(ErrReader(%q)) = _ %v", errReader, err)
	}
}

func TestCountBytes(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"", 0},
		{"one", 3},
		{"one ", 4},
		{"one two", 7},
		{"one two ä", 10},
	}
	for _, test := range tests {
		r := strings.NewReader(test.input)
		if got, err := CountBytes(r); got != test.want || err != nil {
			t.Errorf("CountBytes(%q) = %v %v, want %v", test.input, got, err, test.want)
		}
	}
}

func TestCountBytesErrReader(t *testing.T) {
	errReader := errors.New("reader error")
	r := iotest.ErrReader(errReader)
	if _, err := CountBytes(r); !errors.Is(err, errReader) {
		t.Errorf("CountBytes(ErrReader(%q)) = _ %v", errReader, err)
	}
}
