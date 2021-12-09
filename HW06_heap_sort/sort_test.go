package HW06_heap_sort

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

type SortFunc struct {
	Name string
	Func func([]int)
}

var sortFuncs = []SortFunc{
	{"BubbleSort", BubbleSort},
	{"Insertion ", InsertionSort},
	{"ShellSort ", ShellSort},
	{"Sedgewick ", ShellSortSedgewick},
	{"ShellFrank", ShellSortFrank},
	{"HeapSort  ", HeapSort},
	{"QuickSort ", sort.Ints},
}

var cases = []struct {
	in   []int
	want []int
}{
	{[]int{}, []int{}},
	{[]int{0}, []int{0}},
	{[]int{0, 1}, []int{0, 1}},
	{[]int{1, 0}, []int{0, 1}},
	{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{9, 10, 3, 7, 6, 5, 4, 8, 2, 0, 1}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
}

func TestSort(t *testing.T) {
	for _, f := range sortFuncs {
		for _, c := range cases {
			arr := make([]int, len(c.in))
			copy(arr, c.in)
			f.Func(arr)
			require.Equal(t, c.want, arr)
		}
	}
}

func TestSortRandom(t *testing.T) {
	n := 100
	for _, f := range sortFuncs {
		arr1, arr2 := random(n), random(n)
		require.Equal(t, arr2, arr1)
		f.Func(arr1)
		require.NotEqual(t, arr2, arr1)
		sort.Ints(arr2)
		require.Equal(t, arr2, arr1)
	}
}
