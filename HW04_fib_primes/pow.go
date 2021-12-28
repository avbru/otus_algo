package HW04_fib_primes

func PowSimple(num float64, pow int) float64 {
	res := float64(1)
	for i := 0; i < pow; i++ {
		res = res * num
	}
	return res
}

func PowAdd(num float64, pow int) float64 {
	if pow == 0 {
		return 1.0
	}

	curPow := 1
	delta := 0
	number := num

	for {
		if curPow*curPow > pow {
			delta = pow - curPow
			break
		}

		curPow = curPow * 2
		if curPow == pow {
			return num * num
		}
		num = num * num
	}

	res := num
	for i := 1; i <= delta; i++ {
		res = res * number
	}

	return res
}

func PowN(num float64, pow int) float64 {
	a, p := num, 1.0
	for i := pow / 2; i >= 1; i = i / 2 {
		a = a * a
		if i%2 == 1 {
			p = p * a
		}
	}

	if pow%2 != 0 {
		return p * num
	}
	return p
}
