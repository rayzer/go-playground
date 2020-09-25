package main

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
//Use stack
func largestRectangleArea(heights []int) int {
	maxArea := 0
	var stack []int
	stack = append(stack, -1)

	for i := 0; i < len(heights); i++ {

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
