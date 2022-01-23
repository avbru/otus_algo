package HW18_dt

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTree(t *testing.T) {
	f := [][]int{
		{1},
		{4, 3},
		{7, 5, 9},
		{2, 4, 1, 7},
		{3, 2, 7, 0, 8},
	}
	require.Equal(t, 28, tree(f))
}
