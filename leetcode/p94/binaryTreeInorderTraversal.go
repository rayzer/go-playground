package main

/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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

//0 ms, faster than 100.00%
/*
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	if root.Left != nil {
		tmp := inorderTraversal(root.Left)
		result = append(tmp, result...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		tmp := inorderTraversal(root.Right)
		result = append(result, tmp...)
	}
	return result
}*/
//2: 迭代遍历
func inorderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		//move to left most
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		//pop one and check
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)

		//move right if available
		root = root.Right
	}
	return result
}

// @lc code=end
