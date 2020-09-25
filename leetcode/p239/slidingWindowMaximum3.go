package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */

// @lc code=start
//beautiful solution

func maxSlidingWindow2(nums []int, k int) []int {
	var result, stack []int
	pushBack := func(a int) {
		for idx, val := range stack {
			if val < a {
				stack = stack[:idx]
				break
			}
		}
		stack = append(stack, a)
	}

	for idx, val := range nums {
		pushBack(val)
		if idx < k-1 {
			continue
		}
		result = append(result, stack[0])
		if stack[0] == nums[idx-k+1] {
			stack = stack[1:]
		}
	}
	return result
}

//func maxSlidingWindow(nums []int, k int) []int {
//	var res, stack []int
//	pushBack := func(a int) {
//		for idx, val := range stack {
//			if val < a {
//				stack = stack[:idx] //if value < input value, pop back stack
//				break
//			}
//		}
//		stack = append(stack, a)
//	}
//
//	for idx, val := range nums {
//		pushBack(val)
//		if idx < k-1 { //fill in first window elements first
//			continue
//		}
//		res = append(res, stack[0])
//		if stack[0] == nums[idx-k+1] { //if stack top value == input, the windows is moving out, pop first stack
//			stack = stack[1:]
//		}
//	}
//	return res
//}

// @lc code=end

func main() {
	// input := []int{1, 3, -1, -3, 5, 3, 6, 7}
	// k := 3
	// input := []int{1}
	// k := 1
	input := []int{1, 3, 1, 2, 0, 5}
	k := 3
	fmt.Println(maxSlidingWindow2(input, k))
}
