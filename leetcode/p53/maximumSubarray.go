package main

import "fmt"

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + max(0, nums[i-1])
	}
	return maxN(nums)
}

func maxN(nums []int) int {
	var m int
	for i, e := range nums {
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(input))
}
