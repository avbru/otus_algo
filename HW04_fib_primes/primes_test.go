package HW04_fib_primes

import (
	"fmt"
	"github.com/avbru/algo/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPrimes(t *testing.T) {
	reader := utils.NewReader(t, "5.Primes")

	for reader.Next(t) {
		n := reader.In.Int()
		out := reader.Out.Int()

		now := time.Now()
		res := Primes(n)
		dur := time.Since(now)
		fmt.Printf("case %2d: N: %9d want: %10d have: %10d spent: %v\n", reader.Idx-1, n, out, res, dur)
		if res == -1 {
			println("timeout")
			return
		}
		require.Equal(t, out, res)
	}
}
