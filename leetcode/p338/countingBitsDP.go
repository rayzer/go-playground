package main

import "fmt"

/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */

// Bit + DP
// @lc code=start
func countBitsDP(num int) []int {
	result := make([]int, num+1)
	for i := 1; i <= num; i++ {
		result[i] += result[i&(i-1)] + 1
	}
	return result
	//0 result[1&0 = 0 ]+ 1 = 1
	//0 1 result[2&1 = 0] + 1 = 1
	//0 1 1 result[3&2 = 2] + 1 = 2
	//4 & 3 = 0
	//5 & 4 = 4
	// 6 & 5 = 4

	// res := make([]int, num+1)
	// res[0] = 0
	// if num > 0 {
	// 	res[1] = 1
	// }
	// for i := 2; i <= num; i++ {
	// 	res[i] = res[i/2] + res[i%2]
	// }
	// return res
}

// @lc code=end
func main() {
	fmt.Println(countBitsDP(2))
	fmt.Println(countBitsDP(5))
}
