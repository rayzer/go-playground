package main

/*
 * @lc app=leetcode id=589 lang=golang
 *
 * [589] N-ary Tree Preorder Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}
	result = append(result, root.Val)
	for _, node := range root.Children {
		result = append(result, preorder(node)...)
	}
	return result
}

// @lc code=end
