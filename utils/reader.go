package utils

import (
	"fmt"
	"io"
	"os"
	"testing"
)

type Reader struct {
	In, Out Data
	Idx     int
	path    string
}

func NewReader(t *testing.T, path string) *Reader {
	t.Helper()
	_, err := os.ReadDir(path)
	if err != nil {
		t.Fatal(err)
	}

	return &Reader{path: path}
}

func (r *Reader) Next(t *testing.T) bool {
	t.Helper()
	f, err := os.Open(fmt.Sprintf("%s/test.%d.in", r.path, r.Idx))
	if err != nil {
		fmt.Printf("%s\n", err)
		return false
	}

	r.In, err = io.ReadAll(f)
	if err != nil {
		fmt.Printf("%s", err)
		return false
	}
	f.Close()

	f, err = os.Open(fmt.Sprintf("%s/test.%d.out", r.path, r.Idx))
	if err != nil {
		fmt.Printf("%s", err)
		return false
	}
	r.Out, err = io.ReadAll(f)
	if err != nil {
		fmt.Printf("%s", err)
		return false
	}
	f.Close()
	r.Idx++
	return true
}
