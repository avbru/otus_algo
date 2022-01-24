package HW04_fib_primes

import (
	"fmt"
	"github.com/avbru/algo/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFibGolden(t *testing.T) {
	reader := utils.NewReader(t, "4.Fibo")
	for reader.Next(t) {
		n := reader.In.Int()
		out := reader.Out.String()
		fmt.Println(n)
		res := FibGolden(n)
		require.Equal(t, out, res, n)

	}
}
