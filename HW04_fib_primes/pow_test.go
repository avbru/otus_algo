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
}

func TestPowN(t *testing.T) {
	//require.Equal(t, 2.0, PowN(2, 1))
	//require.Equal(t, math.Pow(2.0, 100.0), PowN(2, 100))
	require.Equal(t, math.Pow(2.0, 51.0), PowN(2, 51))

	//require.Equal(t, 4.0, PowN(2, 2))
	////	require.Equal(t, 8.0, PowN(2, 3))
	//require.Equal(t, 16.0, PowN(2, 4))
	//require.Equal(t, 64.0, PowN(2, 6))
	//require.Equal(t, 128.0, PowN(2, 7))
	//
	//
	///PowPlus(2, 2) //1
	//PowN(2, 2) //4
	//PowN(2, 3) //8
	//PowN(2, 4) //16
	//PowN(2, 5) //32
}
