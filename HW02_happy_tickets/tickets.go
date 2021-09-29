package HW02_happy_tickets

func Tickets(n int) int {
	var vsum []int
	h := []int{1}

	for N := 1; N <= n; N++ {
		l := N*10 - N + 1
		vsum = make([]int, l, l)
		for vi := 0; vi < len(vsum); vi++ {
			sum := 0
			for i := 0; i < 10; i++ {
				idx := vi - i
				if idx < 0 || idx > len(h)-1 {
					continue
				}
				sum += h[idx]
			}
			vsum[vi] = sum
		}
		h = vsum
	}

	total := 0
	for _, v := range vsum {
		total += v * v
	}

	return total
}
