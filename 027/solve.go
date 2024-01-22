func removeElement(nums []int, val int) int {
	// Use two pointers (`off` and `i`),
	// where `off` is the number of elements that are not equal to `val`,
	off := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[off] = nums[i]
			off++
		}
	}

	return off
}