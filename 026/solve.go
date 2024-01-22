func removeDuplicates(nums []int) int {
	// According to the problem, the array is sorted,
	// so we can use two pointers (`start` and `i`)
	prev := nums[0]
	start := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			prev = nums[i]
			nums[i], nums[start] = nums[start], nums[i]
			start++
		}
	}
	return start
}