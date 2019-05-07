func moveZeroes(nums []int) {
	zcnt := 0
	idx := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zcnt++
		} else {
			nums[idx] = nums[i]
			idx++
		}
	}

	for i := idx; i < idx+zcnt; i++ {
		nums[i] = 0
	}
}
