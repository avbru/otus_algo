package HW18_dt

import "fmt"

func tree(f [][]int) int {
	for n := len(f) - 2; n >= 0; n-- {
		for i := 0; i < len(f[n]); i++ {
			if f[n+1][i] > f[n+1][i+1] {
				f[n][i] = f[n][i] + f[n+1][i]
			} else {
				f[n][i] = f[n][i] + f[n+1][i+1]
			}
		}
	}
	fmt.Println(f)
	return f[0][0]
}
