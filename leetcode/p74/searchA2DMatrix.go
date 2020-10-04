package main

import "fmt"

/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) != 0 || (len(matrix) == 1 && len(matrix[0]) == 0) {
		return false
	}
	for i := 0; i <= len(matrix)-1; i++ {
		col := len(matrix[i]) - 1
		if matrix[i][0] <= target && target <= matrix[i][col] {
			return binSearch(matrix[i], target)
		}
	}
	return false
}

func binSearch(nums []int, target int) bool {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if target == nums[mid] {
			return true
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

// @lc code=end

func main() {
	// matrix := [][]int{
	// 	{1, 3, 5, 7},
	// 	{10, 11, 16, 20},
	// 	{23, 30, 34, 50},
	// }
	matrix := [][]int{{}}
	target := 1
	fmt.Println(searchMatrix(matrix, target))

	//test := []int{1, 3, 5, 7}
	//fmt.Println(binSearch(test, 4))
}
