package HW13_min_tree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKruskalDFS(t *testing.T) {
	g := Graph{Edges: edges}
	want := [][3]int{
		{5, 0, 3},
		{5, 2, 4},
		{6, 3, 5},
		{7, 0, 1},
		{7, 1, 4},
		{9, 4, 6},
	}
	require.Equal(t, want, g.KruskalDFS())
}
