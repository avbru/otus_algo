package HW16_kmp

func CreatePiSlow(s string) []int {
	pi := make([]int, len(s)+1)
	for q := 1; q <= len(s); q++ {
		line := s[:q]

		for ln := 1; ln < q; ln++ {
			if line[:ln] == line[len(line)-ln:] {
				pi[q] = ln
			}
		}
	}
	return pi
}

func CreatePi(s string) []int {
	pi := make([]int, len(s)+1)

	for q := 1; q < len(s); q++ {
		ln := pi[q]
		for ln > 0 && (s[ln] != s[q]) && ln <= len(s) {

			ln = pi[ln]
		}

		if s[ln] == s[q] {
			ln++
		}
		pi[q+1] = ln
	}

	return pi
}

// KMP - возвращает слайс индексов вхождений шаблона
func KMP(s, pattern string) []int {
	line := pattern + string([]byte{0}) + s
	pi := CreatePi(line)
	var idx []int
	for k, v := range pi {
		if v == len(pattern) {
			idx = append(idx, k-2*len(pattern)-1)
		}
	}
	return idx
}
