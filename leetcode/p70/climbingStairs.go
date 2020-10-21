package main

import "fmt"

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start

//1: 1
//2: 2
//f(3) = f(2) + f(1)
//f(n) = f(n-1) + f(n-2)

// func climbStairs(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}
// 	return climbStairs(n-1) + climbStairs(n-2)
// }

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a := 1
	b := 2
	result := 0
	for i := 3; i <= n; i++ {
		result = a + b
		a, b = b, result
	}
	return result
}

//TODO:
//变化1：每步可以走1，2，3级台阶
//f(1) = 1
//f(2) = 1-1, 2
//f(3) = 1-2, 1-1-1, 2-1, 3
//f(4) = 1-2-1, 1-1-1, 2-1-1, 3-1
//       1-1-2, 2-2,
//       1-3

//f(n) = f(n-1) + f(n-2) + f(n-3)
//变化2：相邻两部的步伐不能相同
//f(1) = 1
//f(2) = 2
//f(3) = 1-2, 2-1
//f(4) = 1-2-1
//f(5) = 2-1-2
//f(6) = 2-1-2-1, 1-2-1-2
//f(7) = 1-2-1-2-1

// @lc code=end

func main() {
	fmt.Println(climbStairs(10))
}
