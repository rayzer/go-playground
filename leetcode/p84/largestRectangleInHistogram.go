package p84

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */

// @lc code=start
// 暴力法1：
// 双指针遍历，求得每一种组合下的面积
//Exceed time limi
func largestRectangleArea(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		for j := i; j < len(heights); j++ {
			area := (j - i + 1) * min(heights[i:j+1])
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func min(heights []int) int {
	mini := math.MaxInt64
	for i := 0; i < len(heights); i++ {
		if heights[i] < mini {
			mini = heights[i]
		}
	}
	return mini
}

// @lc code=end

func main() {
	input := []int{1, 1}
	fmt.Println(largestRectangleArea(input))
}
