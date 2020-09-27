package main

/*
 * @lc app=leetcode id=590 lang=golang
 *
 * [590] N-ary Tree Postorder Traversal
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

func postorder(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}
	for _, child := range root.Children {
		result = append(result, postorder(child)...)
	}
	return append(result, root.Val)
}

// @lc code=end
