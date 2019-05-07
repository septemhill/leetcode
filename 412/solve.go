func fizzBuzz(n int) []string {
	var s []string

	for i := 1; i <= n; i++ {
		var tmp string
		if i%3 == 0 {
			tmp += "Fizz"
		}

		if i%5 == 0 {
			tmp += "Buzz"
		}

		if i%3 != 0 && i%5 != 0 {
			tmp += strconv.FormatInt(int64(i), 10)
		}
		s = append(s, tmp)
	}
	return s
}
