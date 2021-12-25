package HW16_kmp

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateDelta(t *testing.T) {
	want := [][]int{{1, 0, 0, 0}, {1, 2, 0, 0}, {3, 0, 0, 0}, {1, 4, 0, 0}, {3, 0, 5, 0}}
	require.Equal(t, want, CreateDelta("ABABC"))
}

func TestSearch(t *testing.T) {
	require.Equal(t, 2, Search("ABABABC", "ABABC"))
	require.Equal(t, 8, Search("ABABABC ABABCTDD", "ABABCT"))
}

func Test_uniques(t *testing.T) {
	require.Empty(t, uniques(""))
	require.Equal(t, []rune{'a'}, uniques("a"))
	require.Equal(t, []rune{'a'}, uniques("aa"))
	require.Equal(t, []rune{'a', 'b', 'c', 'd'}, uniques("aabbbccdd"))
}
