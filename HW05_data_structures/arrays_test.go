package HW05_data_structures

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var resizeFuncs = []func(*[]int){resizeSA, resizeVector, resizeFactor}

func TestArrayAdd(t *testing.T) {
	//TEST ADD
	for _, f := range resizeFuncs {
		arr := Array{
			arr:        nil,
			resizeFunc: f,
		}
		for i := 0; i < 12; i++ {
			arr.Append(i)
		}
		require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, arr.arr)
	}

	//TEST PUT
	for _, f := range resizeFuncs {
		arr := Array{
			arr:        nil,
			resizeFunc: f,
		}
		for i := 0; i < 12; i++ {
			arr.Add(i, 0)
		}
		require.Equal(t, []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, arr.arr)
	}

}

func Benchmark_APP_SA__10(b *testing.B) { benchmarkAppend(b, resizeSA, 10) }
func Benchmark_APP_Vec_10(b *testing.B) { benchmarkAppend(b, resizeVector, 10) }
func Benchmark_APP_Fac_10(b *testing.B) { benchmarkAppend(b, resizeFactor, 10) }

func Benchmark_APP_SA__100(b *testing.B) { benchmarkAppend(b, resizeSA, 100) }
func Benchmark_APP_Vec_100(b *testing.B) { benchmarkAppend(b, resizeVector, 100) }
func Benchmark_APP_Fac_100(b *testing.B) { benchmarkAppend(b, resizeFactor, 100) }

func Benchmark_APP_SA__1000(b *testing.B) { benchmarkAppend(b, resizeSA, 1000) }
func Benchmark_APP_Vec_1000(b *testing.B) { benchmarkAppend(b, resizeVector, 1000) }
func Benchmark_APP_Fac_1000(b *testing.B) { benchmarkAppend(b, resizeFactor, 1000) }

func Benchmark_APP_SA__10000(b *testing.B) { benchmarkAppend(b, resizeSA, 10000) }
func Benchmark_APP_Vec_10000(b *testing.B) { benchmarkAppend(b, resizeVector, 10000) }
func Benchmark_APP_Fac_10000(b *testing.B) { benchmarkAppend(b, resizeFactor, 10000) }

func Benchmark_ADD_SA__10(b *testing.B) { benchmarkAdd(b, resizeSA, 10) }
func Benchmark_ADD_Vec_10(b *testing.B) { benchmarkAdd(b, resizeVector, 10) }
func Benchmark_ADD_Fac_10(b *testing.B) { benchmarkAdd(b, resizeFactor, 10) }

func Benchmark_ADD_SA__100(b *testing.B) { benchmarkAdd(b, resizeSA, 100) }
func Benchmark_ADD_Vec_100(b *testing.B) { benchmarkAdd(b, resizeVector, 100) }
func Benchmark_ADD_Fac_100(b *testing.B) { benchmarkAdd(b, resizeFactor, 100) }

func Benchmark_ADD_SA__1000(b *testing.B) { benchmarkAdd(b, resizeSA, 1000) }
func Benchmark_ADD_Vec_1000(b *testing.B) { benchmarkAdd(b, resizeVector, 1000) }
func Benchmark_ADD_Fac_1000(b *testing.B) { benchmarkAdd(b, resizeFactor, 1000) }

func Benchmark_ADD_SA__10000(b *testing.B) { benchmarkAdd(b, resizeSA, 10000) }
func Benchmark_ADD_Vec_10000(b *testing.B) { benchmarkAdd(b, resizeVector, 10000) }
func Benchmark_ADD_Fac_10000(b *testing.B) { benchmarkAdd(b, resizeFactor, 10000) }

func benchmarkAppend(b *testing.B, f func(*[]int), N int) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		arr := Array{
			arr:        nil,
			resizeFunc: f,
		}
		for i := 0; i < N; i++ {
			arr.Append(i)
		}
	}
}

func benchmarkAdd(b *testing.B, f func(*[]int), N int) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		arr := Array{
			arr:        nil,
			resizeFunc: f,
		}
		for i := 0; i < N; i++ {
			arr.Add(i, 0)
		}
	}
}