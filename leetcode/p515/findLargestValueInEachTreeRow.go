package main

import "math"

/*
 * @lc app=leetcode id=515 lang=golang
 *
 * [515] Find Largest Value in Each Tree Row
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var queue []*TreeNode
	var result []int
	var levelResult int
	queue = append(queue, root)
	for len(queue) != 0 {
		count := len(queue)
		levelResult = math.MinInt64
		for _, node := range queue {
			if node.Val > levelResult {
				levelResult = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, levelResult)
		queue = queue[count:]
	}
	return result
}

// @lc code=end
