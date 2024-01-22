func isPalindrome(x int) bool {
	// According to the problem, negative numbers are not palindromes
	if x < 0 {
		return false
	}

	// If the reversed number is the same as the original number, it is a palindrome
	reverse := 0
	for copy := x; copy > 0; copy /= 10 {
		reverse = reverse*10 + copy%10
	}

	return reverse == x
}
