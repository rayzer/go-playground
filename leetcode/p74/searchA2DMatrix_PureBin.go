package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	var m int
	if n > 0 {
		m = len(matrix[0])
	} else {
		return false
	}

	left, right := 0, n*m-1

	for left <= right {
		mid := left + (right-left)/2
		if target == matrix[mid/m][mid%m] {
			return true
		} else if target > matrix[mid/m][mid%m] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// @lc code=end

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	//matrix := [][]int{{}}
	target := 1
	fmt.Println(searchMatrix(matrix, target))

	//test := []int{1, 3, 5, 7}
	//fmt.Println(binSearch(test, 4))
}
