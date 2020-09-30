package main

import (
	google_pubsub_loadtest "cloud.google.com/go/pubsub/loadtest/pb"
	"math"
)

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

func isValidBST(root *TreeNode) bool {
	var stack []*TreeNode
	current := math.MinInt64
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//中序遍历，如果前一个值小于当先值，则不是二叉搜索树
		if root.Val <= current {
			return false
		}
		current = root.Val

		root = root.Right
	}
	return true
}

// @lc code=end
