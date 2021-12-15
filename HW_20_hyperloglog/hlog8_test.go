package hyperloglog

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func TestHLog(t *testing.T) {
	require.Equal(t, 18.799227586206896, HLog8())
}

func TestRandom(t *testing.T) {
	uniques := make(map[uint32]struct{})
	max := (uint32(1) << 24)
	println(max)
	for i := uint32(0); i < max; i++ {
		rnd := rand.Uint32()
		uniques[rnd] = struct{}{}
		if i%10_000_000 == 0 {
			println(i)
		}
	}
	println(2<<32, len(uniques), int(max)-len(uniques))
}
