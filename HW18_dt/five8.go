package HW18_dt

func five8(n int) int {
	f5, f8 := 1, 1
	var f55, f88, n5, n8, n55, n88 int
	for j := 2; j <= n; j++ {
		n5 = f5 + f55
		n8 = f8 + f88
		n55, n88 = f5, f8
		f5, f8 = n5, n8
		f55, f88 = n55, n88
	}
	return f5 + f8 + f55 + f88
}
