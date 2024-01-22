func reverse(x int) int {
	digit := make([]int, 0)

	for i := x; i != 0; {
		digit = append(digit, i%10)
		i = i / 10
	}

	res := 0

	for i := 0; i < len(digit); i++ {
		res = res*10 + digit[i]
	}

	if res > 1<<31-1 || res < -(1<<31) {
		return 0
	}

	return res
}