package main

import "fmt"

/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	var result []string
	_generateParenthesis(0, 0, n, "", &result)
	return result
}

func _generateParenthesis(left int, right int, n int, s string, result *[]string) {
	if left == n && right == n {
		*result = append(*result, s)
		return
	}
	if left < n {
		_generateParenthesis(left+1, right, n, s+"(", result)
	}
	if right < left {
		_generateParenthesis(left, right+1, n, s+")", result)
	}
}

// @lc code=end

func main() {
	fmt.Println(generateParenthesis(3))
}
