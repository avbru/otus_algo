package HW04_fib_primes

import (
	"context"
	"math"
	"time"
)

func Primes(n int) int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := 0
	for j := 2; j <= n; j++ {
		select {
		case <-ctx.Done():
			return -1
		default:
			if IsPrimeBrute(j) {
				q++
			}
		}
	}

	return q
}

func IsPrimeBrute(n int) bool {
	count := 0
	for d := 1; d <= n; d++ {
		if n%d == 0 {
			count++
		}
	}

	return count == 2
}

func IsPrimeBruteEnh(n int) bool {
	for d := 2; d < n; d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
}

func IsPrimeHALF(n int) bool {
	for d := 2; d <= n/2; d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
}

func IsPrimeSQRT(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	for d := 2; d <= sqrt; d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
}

func IsPrimeOdd(n int) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for d := 3; d <= sqrt; d = d + 2 {
		if n%d == 0 {
			return false
		}
	}
	return true
}
