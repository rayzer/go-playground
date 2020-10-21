package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */

// @lc code=start
//Use stack
func largestRectangleArea(heights []int) int {
	heights = append([]int{-2}, heights...) //-2左边界
	heights = append(heights, -1)           //-1右边界
	size := len(heights)
	s := make([]int, 1, size) //初始化len 1， []int{0}，必须要有，见22
	result := 0
	for i := 1; i < size; {
		if heights[s[len(s)-1]] < heights[i] { //栈顶元素小于右边高度，入栈
			s = append(s, i)
			i++
			continue
		}
		fmt.Println("s: ", s, "i", i)
		//栈顶元素大于右边高度，找到右边界，弹出栈顶，计算其对应的最大面积
		//跨度 (i-s[len(s)-2]-1
		// s[len(s)-2] 矩形的左边界
		//栈顶元素高度 heights[s[len(s)-1]]
		result = max(result, heights[s[len(s)-1]]*(i-s[len(s)-2]-1))
		s = s[:len(s)-1]
		fmt.Println("result: ", result)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	input := []int{1, 1}
	fmt.Println(largestRectangleArea(input))
}
