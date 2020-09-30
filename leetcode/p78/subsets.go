package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		return result
	}
	_subsets(&result, nums, []int{}, 0)
	return result
}

func _subsets(result *[][]int, nums []int, sub []int, idx int) {
	if idx == len(nums) {
		*result = append(*result, sub)
		return
	}
	_subsets(result, nums, sub, idx+1)
	newsub := append(sub, nums[idx])
	_subsets(result, nums, newsub, idx+1)
}

// @lc code=end

func main() {
	input := []int{9, 0, 3, 5, 7}
	fmt.Println(subsets(input))
}
