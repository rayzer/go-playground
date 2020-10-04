package main

import "fmt"

/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	endReachable := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i]+i >= endReachable {
			endReachable = i
		}
	}
	return endReachable == 0
}

// @lc code=end
func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}
