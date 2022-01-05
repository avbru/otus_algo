package HW14_dijkstra

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var edges = [][3]int{
	{2, 0, 1},
	{3, 1, 2},
	{6, 1, 4},

	{4, 0, 2},
	{1, 2, 4},

	{9, 0, 3},
	{7, 2, 3},
	{6, 2, 5},
	{4, 4, 5},

	{1, 3, 5},
	{5, 3, 6},
	{8, 5, 6},
}
var T = []string{"B", "A", "C", "E", "D", "F", "G"}

func Test_Dijkstra(t *testing.T) {
	g := Graph{Edges: edges}
	g.Dijkstra(1)

	for k, v := range g.Vertices {
		fmt.Println(T[k], v.W, T[v.From], "used: ", v.Used)
	}
}

func TestCheapest(t *testing.T) {
	g := Graph{Edges: edges}
	g.buildVertices()
	require.Equal(t, 0, g.peekCheapest(1))
	require.Equal(t, 4, g.peekCheapest(2))
}
