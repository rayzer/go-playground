package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	subsets := [][]int{}
	subsets = append(subsets, []int{})
	for _, num := range nums {
		fmt.Println("current sets: ", subsets)
		for _, subset := range subsets {
			fmt.Println("current subsets: ", subset)
			//https://leetcode.com/problems/subsets/discuss/390227/Go-Golang-solution-with-An-Unexpected-Problem-not-seen-in-Java.-With-Detailed-Explanation.
			newSubset := append([]int{}, subset...)
			newSubset = append(newSubset, num)
			subsets = append(subsets, newSubset)
		}
	}
	return subsets
}

// @lc code=end

func main() {
	input := []int{9, 0, 3, 5, 7}
	fmt.Println(subsets(input))
}
