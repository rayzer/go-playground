package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return fastPow(x, n)
}

func fastPow(x float64, n int) (result float64) {
	if n == 0 {
		return 1
	}
	tmp := fastPow(x, n/2)
	if n%2 == 1 {
		result = tmp * tmp * x
	} else {
		result = tmp * tmp
	}
	return
}

// @lc code=end

func main() {
	fmt.Println(myPow(2.0, -2))
}
