package main

import "math"

/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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

func isValidBST(root *TreeNode) bool {
	return _isValidBST(root, math.MaxInt64, math.MinInt64)
}

func _isValidBST(root *TreeNode, upper int, lower int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return _isValidBST(root.Left, root.Val, lower) && _isValidBST(root.Right, upper, root.Val)
}

// @lc code=end
