package HW04_fib_primes

import (
	"math/big"
)

func FibRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibRecursion(n-1) + FibRecursion(n-2)
}

func FibIter(n int) string {
	a := big.NewInt(0)
	b := big.NewInt(1)
	temp := big.NewInt(0)
	for i := 0; i < n; i++ {
		temp = a
		a = b
		b = temp.Add(temp, a)
	}
	return a.String()
}

const prec = uint(7000000)

func FibGolden(n int) string {
	if n < 1000000 {
		return FibIter(n)
	}

	zero := big.NewFloat(0.0).SetPrec(prec)
	one := big.NewFloat(1.0).SetPrec(prec)
	half := big.NewFloat(0.5).SetPrec(prec)
	five := big.NewFloat(5).SetPrec(prec)

	sqrt5 := five.Sqrt(five)
	phi := zero.Add(one, sqrt5)
	phi = phi.Mul(phi, half)

	num := BigPow(phi, n)

	num.Mul(num, sqrt5.Quo(one, sqrt5))
	num.Add(num, half)

	i, _ := num.Int(nil)
	return i.String()
}

func BigPow(num *big.Float, pow int) *big.Float {
	a := num
	p := big.NewFloat(1.0).SetPrec(prec)
	for i := pow / 2; i >= 1; i = i / 2 {
		a.Mul(a, a)
		if i%2 == 1 {
			p.Mul(p, a)
		}
	}

	if pow%2 != 0 {
		return num.Mul(p, num)
	}
	return p
}
