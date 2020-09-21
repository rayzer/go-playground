package main

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */
//遍历求最小差
// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	size := len(nums)
	diff := math.MaxInt32
	for i := 0; i < size && diff != 0; i++ {
		lo := i + 1
		hi := size - 1
		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]
			if abs(target-sum) < abs(diff) {
				diff = target - sum
			}
			if sum < target {
				lo++
			} else {
				hi--
			}
		}
	}
	return target - diff
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// @lc code=end
