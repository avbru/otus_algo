package HW15_bm

func IndexBruteForce2(s, p string) int {
	for is := 0; is <= len(s)-len(p); is++ {
		for ip := 0; ip < len(p); ip++ {
			//fmt.Printf("%d %d %s %s \n", is, ip, string(s[is+ip]), string(p[ip]))
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

