package HW15_bm

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBM(t *testing.T) {
	require.Equal(t, -1, BM("", ""))
	require.Equal(t, -1, BM("AA", "B"))
	require.Equal(t, 0, BM("A", "A"))
	require.Equal(t, 1, BM("BA", "A"))
	require.Equal(t, 0, BM("ABCD", "ABCD"))
	require.Equal(t, 5, BM("AAAAAABCD", "ABCD"))

	fmt.Println(suffixTable("abcababc"))
}

func TestIndexBruteForce(t *testing.T) {
	require.Equal(t, -1, IndexBruteForce("", ""))
	require.Equal(t, -1, IndexBruteForce("AA", "B"))
	require.Equal(t, 0, IndexBruteForce("A", "A"))
	require.Equal(t, 1, IndexBruteForce("BA", "A"))
	require.Equal(t, 0, IndexBruteForce("ABCD", "ABCD"))
	require.Equal(t, 5, IndexBruteForce("AAAAAABCD", "ABCD"))
}
