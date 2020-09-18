package p09

import "strconv"

//4 ms, faster than 99.26%
func isPalindrome(x int) bool {
	text := strconv.Itoa(x)
	i := 0
	j := len(text) - 1
	for i < j {
		if text[i] != text[j] {
			return false
		}
		i++
		j--
	}
	return true
}
