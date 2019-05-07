func isAnagram(s string, t string) bool {
	m := make(map[string]int)

	for _, val := range s {
		m[string(val)]++
	}

	for _, val := range t {
		m[string(val)]--
	}

	for _, val := range m {
		if val != 0 {
			return false
		}
	}

	return true
}
