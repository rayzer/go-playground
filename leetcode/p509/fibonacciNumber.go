package main

/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */

// @lc code=start
func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	a := 0
	b := 1
	result := 0
	for i := 2; i <= N; i++ {
		result = a + b
		a, b = b, result
	}
	return result
}

// @lc code=end
