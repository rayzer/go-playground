package p239

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */

// @lc code=start
//超时
func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	var tmp_max int

	for i := 0; i < len(nums)-(k-1); i++ {
		tmp_max = max(nums[i : i+k])
		result = append(result, tmp_max)
	}
	return result
}

func max(nums []int) int {
	fmt.Println(nums)
	tmp_max := math.MinInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] > tmp_max {
			tmp_max = nums[i]
		}
	}
	return tmp_max
}

// @lc code=end

func main() {
	input := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(input, k))
}
