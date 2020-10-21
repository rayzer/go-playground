package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		fmt.Println(left, mid, right)
		if nums[mid] < nums[right] {
			//右半有序，排除
			right = mid
		} else if nums[mid] > nums[right] {
			//右半无序，继续查
			left = mid + 1
		}
	}
	return nums[left]
}

// @lc code=end

func main() {
	//input := []int{1, 2}
	input := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(input))
}
