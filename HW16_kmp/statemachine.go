package HW16_kmp

type SM struct {
}

func CreateDelta(pattern string) [][]int {
	alphabet := uniques(pattern)
	alphabet = append(alphabet, 0) //Null string == ANY

	delta := make([][]int, len(pattern))
	for k := range delta {
		delta[k] = make([]int, len(alphabet))
	}

	for q := 0; q < len(pattern); q++ {
		for ia, a := range alphabet {
			line := pattern[0:q] + string(a)
			k := q + 1
			for k >= 0 {
				if pattern[0:k] == line[len(line)-k:] {
					break
				}
				k--
			}

			delta[q][ia] = k
		}
	}
	return delta
}

func Search(text, pattern string) int {
	delta := CreateDelta(pattern)
	alphabet := uniques(pattern)
	abc := make(map[rune]int)
	for k, v := range alphabet {
		abc[v] = k
	}

	q := 0

	for i, r := range text {
		q = delta[q][abc[r]]
		if q == len(pattern) {
			return i - len(pattern) + 1
		}
	}
	return -1
}

func uniques(s string) []rune {
	if s == "" {
		return nil
	}

	var uniq []rune
outer:
	for _, r := range s {
		for _, u := range uniq {
			if u == r {
				continue outer
			}
		}
		uniq = append(uniq, r)
	}
	return uniq
}
