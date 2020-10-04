package main

import "fmt"

/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	t := x
	for t*t > x {
		t = (t + x/t) / 2
	}
	return t
}

// @lc code=end

func main() {
	fmt.Println(mySqrt(9))
}
