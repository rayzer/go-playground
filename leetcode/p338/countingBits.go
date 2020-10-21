package main

import "fmt"

/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */

// Bit
// @lc code=start
func countBits(num int) []int {
	var result []int
	for i := 0; i <= num; i++ {
		var count int
		for j := i; j > 0; count++ {
			j &= j - 1
		}
		result = append(result, count)
	}
	return result
}

// @lc code=end
func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
}
