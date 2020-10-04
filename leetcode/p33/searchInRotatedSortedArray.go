package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[0] <= nums[mid] && (target > nums[mid] || target < nums[0]) {
			//左半升序
			//target 不在左半
			left = mid + 1
		} else if target > nums[mid] && target < nums[0] {
			//左半有旋转 [ 7 0 1 2 3 4 5 6 ]
			//右半升序
			//target 在右半
			left = mid + 1
		} else {
			//继续探索左半
			right = mid
		}
	}
	if left == right && nums[right] == target {
		return left
	} else {
		return -1
	}
}

// @lc code=end

func main() {
	input := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(input, target))
}
