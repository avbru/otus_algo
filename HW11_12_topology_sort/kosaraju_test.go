package HW11_12_topology_sort

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

var vec = [][]int{
	{1},       // 0
	{2, 4, 5}, // 1
	{3},       // 2
	{2, 7},    // 3
	{0, 5},    // 4
	{6},       // 5
	{5, 7},    // 6
	{3},       // 7
}

func TestKosaraju(t *testing.T) {
	fmt.Println(Kosaraju(vec))
	res := [][]int{
		{7, 3, 2},
		{5, 6},
		{4, 1, 0},
	}
	require.Equal(t, res, Kosaraju(vec))
}

func TestReverse(t *testing.T) {
	rev := [][]int{
		{4},       // 0
		{0},       // 1
		{1, 3},    // 2
		{2, 7},    // 3
		{1},       // 4
		{1, 4, 6}, // 5
		{5},       // 6
		{3, 6},    // 7
	}
	require.Equal(t, rev, reverse(vec))
}
