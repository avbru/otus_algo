package utils

import (
	"strconv"
	"strings"
)

type Data []byte

type ChessMoves struct {
	N    int
	Mask uint64
}

func (d Data) String() string {
	return strings.TrimSpace(string(d))
}

func (d Data) RawString() string {
	return string(d)
}

func (d Data) Int() int {
	i, err := strconv.Atoi(strings.TrimSpace(string(d)))
	if err != nil {
		return 0
	}
	return i
}

func (d Data) ChessMoves() ChessMoves {
	fields := strings.Fields(string(d))
	n, _ := strconv.Atoi(fields[0])
	mask, _ := strconv.ParseUint(fields[1], 10, 64)
	return ChessMoves{N: n, Mask: mask}
}
