package HW15_bm

import "strings"

func BM(s, p string) int {
	if len(s) == 0 || len(p) == 0 {
		return -1
	}

	prefix := prefixTable(p)
	suffix := suffixTable(p)
	i := len(p) - 1
	for i < len(s) {
		j := len(p) - 1
		for j >= 0 && s[i] == p[j] {
			i--
			j--
		}
		if j < 0 {
			return i + 1
		}
		i += max(prefix[s[i]], suffix[j])
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func suffixTable(pattern string) []int {
	m := len(pattern)
	table := make([]int, m)

	last := m - 1

	lastPrefix := last
	for i := last; i >= 0; i-- {
		if strings.HasPrefix(pattern, pattern[i+1:]) {
			lastPrefix = i + 1
		}

		table[i] = lastPrefix + last - i
	}

	for i := 0; i < last; i++ {
		lenSuffix := suffixLength(pattern, pattern[1:i+1])
		if pattern[i-lenSuffix] != pattern[last-lenSuffix] {
			table[last-lenSuffix] = lenSuffix + last - i
		}
	}

	return table
}

func prefixTable(p string) [256]int {
	var t [256]int
	for i := range t {
		t[i] = len(p)
	}

	for i := 0; i < len(p)-1; i++ {
		t[p[i]] = len(p) - 1 - i
	}
	return t
}

func suffixLength(a, b string) (i int) {
	for ; i < len(a) && i < len(b); i++ {
		if a[len(a)-1-i] != b[len(b)-1-i] {
			break
		}
	}
	return
}

func IndexBruteForce(s, p string) int {
	for is := 0; is <= len(s)-len(p); is++ {
		for ip := 0; ip < len(p); ip++ {
			if s[is+ip] != p[ip] {
				break
			}
			if ip == len(p)-1 {
				return is
			}
		}
	}

	return -1
}
