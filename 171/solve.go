func titleToNumber(s string) int {
	r := []rune(s)

	sum := 0

	for i := 0; i < len(r); i++ {
		sum += int(r[i]-0x41+1) * int(math.Pow(26, float64(len(r)-i-1)))
	}

	return sum
}
