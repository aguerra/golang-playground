// Package wc implements the core functionality of the traditional wc
// command line tool.
package wc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func count(r io.Reader, f bufio.SplitFunc) (int, error) {
	sc := bufio.NewScanner(r)
	sc.Split(f)
	var count int
	for sc.Scan() {
		count++
	}
	if err := sc.Err(); err != nil {
		return count, fmt.Errorf("count failed: %w", err)
	}
	return count, nil
}

// We adopt the POSIX definition of line: a sequence of zero or more
// non-newline characters plus a terminating newline character.
func scanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		return i + 1, []byte{}, nil
	}
	return 0, nil, nil
}

func CountWords(r io.Reader) (int, error) {
	return count(r, bufio.ScanWords)
}

func CountLines(r io.Reader) (int, error) {
	return count(r, scanLines)
}

func CountBytes(r io.Reader) (int, error) {
	return count(r, bufio.ScanBytes)
}
