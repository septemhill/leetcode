func searchRange(nums []int, target int) []int {
	start, end := -1, -1
	for i, j := 0, len(nums)-1; i <= j; {
		if nums[i] == target && start == -1 {
			start = i
		}

		if nums[j] == target && end == -1 {
			end = j
		}

		if start == -1 {
			i++
		}

		if end == -1 {
			j--
		}

		if start != -1 && end != -1 {
			break
		}
	}

	return []int{start, end}
}