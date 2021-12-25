package HW15_bm

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIndexBruteForce(t *testing.T) {
	require.Equal(t, -1, IndexBruteForce2("", ""))
	require.Equal(t, -1, IndexBruteForce2("AA", "B"))
	require.Equal(t, 0, IndexBruteForce2("A", "A"))
	require.Equal(t, 1, IndexBruteForce2("BA", "A"))
	require.Equal(t, 0, IndexBruteForce2("ABCD", "ABCD"))
	require.Equal(t, 5, IndexBruteForce2("AAAAAABCD", "ABCD"))
}
