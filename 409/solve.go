func longestPalindrome(s string) int {
	m := make(map[byte]int)
	long := 0

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for key := range m {
		if m[key] > 1 {
			long += (m[key] / 2) * 2
			m[key] -= (m[key] / 2) * 2
		}
	}

	for key := range m {
		if m[key] == 1 {
			long++
			break
		}
	}

	return long
}
