package HW06_heap_sort

import (
	"fmt"
	"math/rand"
	"testing"
)

type FillFunc struct {
	Name string
	Func func(int) []int
}

var fillFuncs = []FillFunc{
	{"random ", random},
	{"digits ", digits},
	{"sorted ", sorted},
	{"reverse", reverse},
}

func BenchmarkSort(b *testing.B) {
	b.StopTimer()
	for n := 1000; n <= 100_000; n *= 10 {
		for _, fillFunc := range fillFuncs {
			for _, sortFunc := range sortFuncs {
				b.Run(fmt.Sprintf("N=%d % s %s ", n, fillFunc.Name, sortFunc.Name), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						arr := fillFunc.Func(n)
						b.StartTimer()
						sortFunc.Func(arr)
						b.StopTimer()
					}
				})
			}
		}
	}
}

func random(n int) []int {
	rand.Seed(1)
	arr := make([]int, n, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Int()
	}
	return arr
}

func digits(n int) []int {
	rand.Seed(1)
	arr := make([]int, n, n)
	for i := 0; i < n; i++ {
		arr[i] = int(rand.Uint64()/2-1) % 10
	}
	return arr
}

func sorted(n int) []int {
	rand.Seed(1)
	arr := make([]int, n, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	for i := 1; i < n; i *= 10 {
		arr[i] = rand.Int()
	}

	return arr
}

func reverse(n int) []int {
	arr := make([]int, n, n)
	for i, j := n-1, 0; i >= 0; i, j = i-1, j+1 {
		arr[i] = j
	}
	return arr
}
