package HW10_hash_table

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func TestHashTable(t *testing.T) {
	ht := New()
	ht.Put(Hint(0), "V0")
	ht.Put(Hint(1), "V1")
	ht.Put(Hint(2), "V2")
	ht.Put(Hint(3), "V3")

	ht.Put(Hint(0), "new V0")
	ht.Delete(Hint(0))
	ht.draw()

	ht.Put(Hint(10), "V10")
	ht.Put(Hint(11), "V11")
	ht.Put(Hint(12), "V12")
	ht.Put(Hint(13), "V13")

	ht.Put(Hint(8), "V8")
	v, ok := ht.Get(Hint(11))
	require.True(t, ok)
	require.Equal(t, "V11", v)
	ht.draw()

}

func BenchmarkHashTable(b *testing.B) {
	for n := 1_000; n <= 1_000_000; n *= 10 {
		ht := New()
		rnd := make([]Hint, n)
		//Заполняем на 30%
		for i := 0; i < n; i++ {
			rnd[i] = Hint(rand.Uint32())
			if i%3 == 0 {
				ht.Put(Hint(rand.Uint32()), "")
			}
		}

		b.Run(fmt.Sprintf("GET N=%d ", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for i := 0; i < 1000; i++ {
					ht.Get(rnd[i])
				}
			}
		})

		b.Run(fmt.Sprintf("PUT N=%d ", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for i := 0; i < 1000; i++ {
					ht.Put(rnd[i], "")
				}
			}
		})

		b.Run(fmt.Sprintf("DELETE N=%d ", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for i := 0; i < 1000; i++ {
					ht.Delete(rnd[i])
				}
			}
		})
	}

}
