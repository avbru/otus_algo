package HW18_dt

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIslands(t *testing.T) {
	f := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 1, 1},
	}
	require.Equal(t, 2, islands(f))
}
