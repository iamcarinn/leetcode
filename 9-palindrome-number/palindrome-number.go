func isPalindrome(x int) bool {
	res := false
	cur := x
	reserved := 0
	if cur > 0 {
		for cur > 0 {
			digit := cur % 10
			reserved = reserved*10 + digit
			cur /= 10
		}
	}
    if reserved == x {
		res = true
	}
	return res
}