package HW11_12_topology_sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTarjan(t *testing.T) {
	vec := [][]int{
		{1, 3},    //0
		{2, 4},    //1
		{4, 5},    //2
		{1, 4, 6}, //3
		{7, 8},    //4
		{8},       //5
		{4},       //6
		{},        //7
		{},        //8
		{10},      //9
		{},        //10
	}

	fmt.Println(Tarjan(vec))
	require.Equal(t, []int{9, 10, 0, 3, 6, 1, 2, 5, 4, 8, 7}, Tarjan(vec))
}
