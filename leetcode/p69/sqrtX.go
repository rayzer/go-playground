package main

import "fmt"

/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left, right := 1, x
	for left < right {
		mid := left + (right-left)/2
		if mid > x/mid {
			right = mid
		} else {
			if mid+1 > x/(mid+1) {
				return mid
			}
			left = mid + 1
		}
	}
	return 0
}

// @lc code=end

func main() {
	fmt.Println(mySqrt(1))
}
