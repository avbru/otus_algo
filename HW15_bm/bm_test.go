package HW15_bm

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var cases = []struct {
	want int
	s, p string
}{
	{-1, "", ""},
	{-1, "AA", "B"},
	{0, "A", "A"},
	{1, "BA", "A"},
	{0, "ABCD", "ABCD"},
	{5, "AAAAAABCD", "ABCD"},
}

func TestBM(t *testing.T) {
	for _, c := range cases {
		require.Equal(t, c.want, BM(c.s, c.p))
		require.Equal(t, c.want, IndexBruteForce(c.s, c.p))
		require.Equal(t, c.want, IndexPrefix(c.s, c.p))
		require.Equal(t, c.want, IndexSuffix(c.s, c.p))
	}
}

var benches = []struct {
	s, p string
	name string
}{
	{strings.Repeat("abcdefghijklmnopqrstuvwxy", 100) + "z", "abcdefghijklmnopqrstuvwxyz", "long pattern"},
	{strings.Repeat("a", 1000), "bcdefghijklmnopqrstuvwxyz", "no pattern"},
	{strings.Repeat("abcdefghijklmnopqrstuvwxy", 100) + "z", "z", "short pattern"},
}

func BenchmarkBM(b *testing.B) {
	for _, bn := range benches {
		b.Run(fmt.Sprintf("%s brute force", bn.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IndexBruteForce(bn.s, bn.p)
			}
		})
		b.Run(fmt.Sprintf("%s prefix     ", bn.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IndexPrefix(bn.s, bn.p)
			}
		})
		b.Run(fmt.Sprintf("%s suffix     ", bn.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IndexSuffix(bn.s, bn.p)
			}
		})

		b.Run(fmt.Sprintf("%s boyer      ", bn.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BM(bn.s, bn.p)
			}
		})

	}
}
