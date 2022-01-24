package HW08_bst

import (
	"fmt"
	"math/rand"
	"testing"
)

const N = 1000

func BenchmarkBst(b *testing.B) {
	num := make([]int, N)
	for i := 0; i < N; i++ {
		num[i] = i
	}

	rnd := make([]int, N)
	for i := 0; i < N; i++ {
		rnd[i] = rand.Int()
	}

	sorted := NewBST()
	random := NewBST()

	for _, v := range num {
		sorted.Insert(v)
	}
	for _, v := range rnd {
		random.Insert(v)
	}

	b.Run(fmt.Sprintf("N=%d  %s ", N, "sorted insert"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range num {
				sorted.Insert(v)
			}
		}
	})

	b.Run(fmt.Sprintf("N=%d  %s ", N, "sorted search"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 1; j < N; j *= 10 {
				sorted.Search(num[j])
			}
		}
	})

	b.Run(fmt.Sprintf("N=%d  %s ", N, "sorted delete"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 1; j < N; j = j * 10 {
				sorted.Delete(num[j])
			}
		}
	})

	b.Run(fmt.Sprintf("N=%d  %s ", N, "random insert"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range rnd {
				random.Insert(v)
			}
		}
	})

	b.Run(fmt.Sprintf("N=%d  %s ", N, "random search"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 1; j < N; j *= 10 {
				random.Search(rnd[j])
			}
		}
	})

	b.Run(fmt.Sprintf("N=%d  %s ", N, "random delete"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 1; j < N; j *= 10 {
				random.Delete(rnd[j])
			}
		}
	})
}
