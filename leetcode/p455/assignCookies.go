package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=455 lang=golang
 *
 * [455] Assign Cookies
 */

// @lc code=start
func findContentChildren(greedy []int, size []int) int {
	sort.Ints(greedy)
	sort.Ints(size)
	i, j := 0, 0
	for i < len(greedy) && j < len(size) {
		if greedy[i] <= size[j] {
			i++ //there is one child's greedy satisfied
		}
		j++
	}
	return i
}

// @lc code=end
func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
}
