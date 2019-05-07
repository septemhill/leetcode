func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}

	for key, val := range m {
		if val == 1 {
			return key
		}
	}

	return -1
}
