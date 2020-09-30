package main /*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)

	if leftDepth == 0 {
		return rightDepth + 1
	}
	if rightDepth == 0 {
		return leftDepth + 1
	}
	return min(leftDepth, rightDepth) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
