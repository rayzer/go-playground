package main

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start

func maxArea(height []int) int {
	maxArea := 0
	i := 0
	j := len(height) - 1
	for i < j {
		var area int
		distance := j - i
		var minHeight int
		if height[i] > height[j] {
			minHeight = height[j]
		} else {
			minHeight = height[i]
		}

		area = distance * minHeight
		if area > maxArea {
			maxArea = area
		}

		if height[i] > height[j] {
			j-- ////move from shorter one to seek higher one
		} else {
			i++
		}
	}
	return maxArea
}

// 1124 ms, faster than 5.09%
// func maxArea(height []int) int {
// 	maxArea := 0
// 	for i := 0; i < len(height)-1; i++ {
// 		for j := 1; j < len(height); j++ {
// 			var area int
// 			if height[i] > height[j] {
// 				area = (j - i) * height[j]
// 			} else {
// 				area = (j - i) * height[i]
// 			}
// 			if area > maxArea {
// 				maxArea = area
// 			}
// 		}
// 	}
// 	return maxArea
// }

// @lc code=end
