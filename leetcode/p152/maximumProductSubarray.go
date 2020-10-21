package main

import "fmt"

/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start
func maxProduct(nums []int) int {
	n
	mi, ma, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			mi, ma = ma, mi //记住最正和最负
		}
		ma = max(nums[i], ma*nums[i])
		mi = min(nums[i], mi*nums[i])
		result = max(result, ma)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{-3, -1, -1}))
	fmt.Println(maxProduct([]int{-2, 3, -4}))
}
