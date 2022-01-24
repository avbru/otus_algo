package HW18_dt

func islands(f [][]int) int {
	count := 0
	for i := 0; i < len(f); i++ {
		for j := 0; j < len(f); j++ {
			if f[i][j] == 1 {
				count++
				walk(f, i, j)
			}
		}
	}

	return count
}

func walk(f [][]int, x, y int) {
	if x < 0 || x > len(f)-1 {
		return
	}
	if y < 0 || y > len(f)-1 {
		return
	}

	if f[x][y] == 0 {
		return
	}
	f[x][y] = 0
	walk(f, x-1, y)
	walk(f, x+1, y)
	walk(f, x, y-1)
	walk(f, x, y+1)
}
