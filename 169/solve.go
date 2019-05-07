func majorityElement(nums []int) int {
	m := make(map[int]int)
	max, val := 0, 0

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] > max {
			max = m[nums[i]]
			val = nums[i]
		}
	}

	return val
}
