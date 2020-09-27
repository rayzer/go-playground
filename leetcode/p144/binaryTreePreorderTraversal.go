package main

/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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
	Right *TreeNode
	Left  *TreeNode
}

//func preorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return []int{}
//	}
//	var result []int
//	result = append(result, root.Val)
//	if root.Left != nil {
//		tmp := preorderTraversal(root.Left)
//		result = append(result, tmp...)
//	}
//	if root.Right != nil {
//		tmp := preorderTraversal(root.Right)
//		result = append(result, tmp...)
//	}
//	return result
//}

func preorderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	for root != nil || len(stack) == 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		}
	}

	return result
}

// @lc code=end
