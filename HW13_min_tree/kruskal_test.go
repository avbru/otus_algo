package HW13_min_tree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var edges = [][3]int{
	{5, 0, 3},
	{7, 0, 1},
	{9, 1, 3},
	{8, 1, 2},
	{7, 1, 4},
	{5, 2, 4},
	{15, 3, 4},
	{6, 3, 5},
	{8, 4, 5},
	{9, 4, 6},
	{11, 5, 6},
}

func TestKruskal(t *testing.T) {
	g := Graph{Edges: edges}
	want := [][3]int{
		{5, 0, 3},
		{5, 2, 4},
		{6, 3, 5},
		{7, 0, 1},
		{7, 1, 4},
		{9, 4, 6},
	}

	require.Equal(t, want, g.Kruskal())
}
