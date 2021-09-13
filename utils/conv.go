package utils

import (
	"strconv"
	"strings"
)

type Data []byte

func (d Data) String() string {
	return strings.TrimSpace(string(d))
}

func (d Data) Int() int {
	i, err := strconv.Atoi(strings.TrimSpace(string(d)))
	if err != nil {
		return 0
	}
	return i
}
