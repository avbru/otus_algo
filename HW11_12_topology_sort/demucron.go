package HW11_12_topology_sort

func Demucron(vec [][]int) [][]int {
	mx := vec2Matrix(vec)
	sum := colSum(mx)
	var levels [][]int

	for {
		zeroes := zeroIdx(sum)
		if len(zeroes) == 0 {
			break
		}
		levels = append(levels, []int{})
		for _, v := range zeroes {
			vecSub(sum, mx[v])
			sum[v] = -1
		}

		levels[len(levels)-1] = append(levels[len(levels)-1], zeroes...)
	}

	return levels
}

func colSum(mx [][]int) []int {
	sum := make([]int, len(mx))
	for i := 0; i < len(mx); i++ {
		s := 0
		for j := 0; j < len(mx); j++ {
			s = s + mx[j][i]
		}
		sum[i] = s
	}
	return sum
}

func vecSub(a, b []int) {
	for i := 0; i < len(a); i++ {
		a[i] = a[i] - b[i]
	}
}

func zeroIdx(vec []int) (idx []int) {
	for k, v := range vec {
		if v == 0 {
			idx = append(idx, k)
		}
	}
	return
}

func vec2Matrix(vec [][]int) [][]int {
	matrix := make([][]int, len(vec))
	for k, row := range vec {
		matrix[k] = make([]int, len(vec))
		for _, v := range row {
			matrix[k][v] = 1
		}
	}

	return matrix
}
