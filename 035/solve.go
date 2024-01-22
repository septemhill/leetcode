func searchInsert(nums []int, target int) int {
	pos := 0

	for i := 0; i < len(nums); i++ {
		fmt.Println(target, nums[i])
		if target <= nums[i] {
			pos = i
			return pos
		}
	}

	if pos == 0 {
		return len(nums)
	}
	return pos
}