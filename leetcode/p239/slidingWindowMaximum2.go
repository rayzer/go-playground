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
func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	var stack []int
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			stack = push(stack, nums[i])
		} else {
			stack = push(stack, nums[i])
			result = append(result, max(stack))
			stack = pop(stack, nums[i-k+1])
		}
	}
	return result
}

func max(stack []int) int {
	return stack[0]
}

func push(stack []int, n int) []int {
	fmt.Println(stack, n)
	for len(stack) != 0 && stack[len(stack)-1] < n {
		stack = stack[:len(stack)-1]
	}
	stack = append(stack, n)
	fmt.Println(stack)
	return stack
}

func pop(stack []int, n int) []int {
	fmt.Println("pop", stack, n)
	if len(stack) != 0 && stack[0] == n {
		if len(stack) > 1 {
			stack = stack[1:]
		} else {
			stack = []int{}
		}
	}
	fmt.Println("pop after", stack)
	return stack
}

// func maxSlidingWindow(nums []int, k int) []int {
// 	var result []int
// 	var stack []int
// 	var max_idx int

// 	if k == 1 {
// 		return nums
// 	}

// 	for i := 0; i < k; i++ {
// 		cleanDeque(stack, nums, i, k)
// 		stack = append(stack, i)
// 		if nums[i] > nums[max_idx] {
// 			max_idx = i
// 		}
// 	}
// 	result = append(result, nums[max_idx])

// 	for i := k; i < len(nums); i++ {
// 		cleanDeque(stack, nums, i, k)
// 		stack = append(stack, i)
// 		result = append(result, nums[stack[0]])
// 	}

// 	return result
// }

// func cleanDeque(stack []int, nums []int, i int, k int) {
// 	fmt.Println(stack)
// 	if len(stack) != 0 && stack[0] == i-k {
// 		stack = stack[1:len(stack)]
// 	}
// 	for len(stack) != 0 && nums[i] > nums[stack[len(stack)-1]] {
// 		stack = stack[:len(stack)-1]
// 	}
// }

// @lc code=end

func main() {
	// input := []int{1, 3, -1, -3, 5, 3, 6, 7}
	// k := 3
	// input := []int{1}
	// k := 1
	input := []int{1, 3, 1, 2, 0, 5}
	k := 3
	fmt.Println(maxSlidingWindow(input, k))
}
