package HW04_fib_primes

import (
	"fmt"
	"github.com/avbru/algo/utils"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestPow(t *testing.T) {
	require.Equal(t, float64(8), PowSimple(2, 3))

	reader := utils.NewReader(t, "3.Power")

	for reader.Next(t) {
		num, pow := reader.In.FloatInt()
		out := reader.Out.Float()

		res := PowN(num, pow)
		delta := 0.0000001
		fmt.Printf("case %d: want: %2f have: %2f PASS: %v\n", reader.Idx-1, out, res, (res-out) < delta)
		require.Equal(t, math.Abs(res-out) < delta, true)
	}

	reader = utils.NewReader(t, "3.Power")
	for reader.Next(t) {
		num, pow := reader.In.FloatInt()
		out := reader.Out.Float()

		res := PowAdd(num, pow)
		delta := 0.0000001
		fmt.Printf("case %d: want: %2f have: %2f PASS: %v\n", reader.Idx-1, out, res, (res-out) < delta)
		require.Equal(t, math.Abs(res-out) < delta, true)
	}
}
