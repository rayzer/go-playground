package main

import "fmt"

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
//Runtime: 0 ms, faster than 100.00%
func isValid(s string) bool {
	ordered := []byte(s)
	stack := []byte{}
	mapping := map[byte]byte{']': '[', '}': '{', ')': '('}
	for i := 0; i < len(ordered); i++ {
		t := ordered[i]
		if t == '[' || t == '{' || t == '(' {
			stack = append(stack, t)
		} else {
			if len(stack) == 0 {
				return false
			}
			p := stack[len(stack)-1]
			if mapping[t] == p {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

// @lc code=end

func main() {
	fmt.Println(isValid("(])"))
}
