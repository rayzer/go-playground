package main

import "fmt"

/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 */

//有且仅有一个位为1
// @lc code=start
func isPowerOfTwo(n int) bool {
	return n != 0 && (n&(n-1)) == 0
}

// @lc code=end

func main() {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(3))
	fmt.Println(isPowerOfTwo(4))
	fmt.Println(isPowerOfTwo(5))
}
