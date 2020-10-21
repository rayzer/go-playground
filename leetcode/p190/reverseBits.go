package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	result := uint32(0)
	for i := 0; i < 32; i++ {
		result = (result << 1) + (num & 1)
		num = num >> 1
	}
	return result
}

// @lc code=end
func main() {
	input, err := strconv.ParseUint("00000010100101000001111010011100", 2, 32)
	if err != nil {
		panic(err)
	}

	fmt.Println(reverseBits(uint32(input))) //964176192
}
