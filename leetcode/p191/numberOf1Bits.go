package main

import "fmt"

/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
func hammingWeight(num uint32) int {
	var count int
	for ; num > 0; count++ {
		num &= num - 1 //每次去掉一个1的位
	}
	return count
}

// @lc code=end
func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
	fmt.Println(hammingWeight(00000000000000000000000010000000))
}
