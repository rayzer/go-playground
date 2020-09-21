package main

import "fmt"

/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */

// @lc code=start
//0 ms, faster than 100.00%
func longestValidParentheses(s string) int {
	ordered := []byte(s)
	stack := []int{-1} //?
	var maxLength int
	for i := 0; i < len(ordered); i++ {
		if ordered[i] == '(' {
			stack = append(stack, i)
		} else { // == ')'
			stack = stack[:len(stack)-1]
			fmt.Println(stack)
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				p := stack[len(stack)-1]
				if maxLength < i-p {
					maxLength = i - p
				}
			}
		}
	}
	return maxLength
}

// @lc code=end

func main() {
	fmt.Println(longestValidParentheses(")))"))
}
