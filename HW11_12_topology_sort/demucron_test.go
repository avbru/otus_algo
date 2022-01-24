package HW11_12_topology_sort

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDemucron(t *testing.T) {
	vec := [][]int{
		{2, 12}, //V1
		{12},
		{},
		{2},
		{2, 8, 9},
		{3, 10, 11, 12}, //V6
		{5},
		{1, 3, 5, 13}, //V8
		{0, 6},
		{0, 11, 13},
		{2}, //V11
		{},
		{2},
		{10},
	}

	levels := [][]int{
		{4, 7},
		{1, 8, 9},
		{0, 6, 13},
		{5},
		{3, 10, 11, 12},
		{2},
	}
	printMatrix(vec2Matrix(vec))
	println("-----Levels ----------------")
	printMatrix(Demucron(vec))
	require.Equal(t, levels, Demucron(vec))

}

func printMatrix(m [][]int) {
	for _, row := range m {
		fmt.Println(row)
	}
}
