package HW04_fib_primes

import (
	"fmt"
	"github.com/avbru/algo/utils"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
	"time"
)

func TestPow(t *testing.T) {
	require.Equal(t, float64(8), PowSimple(2, 3))

	reader := utils.NewReader(t, "3.Power")

	for reader.Next(t) {
		num, pow := reader.In.FloatInt()
		out := reader.Out.Float()
		now := time.Now()
		res := PowN(num, pow)
		dur := time.Since(now)
		delta := 0.0000001
		fmt.Printf("case %d: want: %2f have: %2f PASS: %v spent: %s \n", reader.Idx-1, out, res, (res-out) < delta, dur)
		require.Equal(t, math.Abs(res-out) < delta, true)
	}

	reader = utils.NewReader(t, "3.Power")
	for reader.Next(t) {
		num, pow := reader.In.FloatInt()
		out := reader.Out.Float()
		now := time.Now()
		res := PowAdd(num, pow)
		dur := time.Since(now)
		delta := 0.0000001
		fmt.Printf("case %d: want: %2f have: %2f PASS: %v spent: %s\n", reader.Idx-1, out, res, (res-out) < delta, dur)
		require.Equal(t, math.Abs(res-out) < delta, true)
	}
}
