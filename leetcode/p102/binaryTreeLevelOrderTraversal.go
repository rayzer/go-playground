package main

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */

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

// @lc code=start

func levelOrder(root *TreeNode) [][]int {
	var queue []*TreeNode
	result := make([][]int, 0)
	level := 0
	var BFS func(root *TreeNode)
	BFS = func(root *TreeNode) {
		queue = append(queue, root)
		for len(queue) != 0 {
			count := len(queue)
			result = append(result, []int{})
			for _, node := range queue {
				result[level] = append(result[level], node.Val)
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
			queue = queue[count:]
			level++
		}
	}
	BFS(root)
	return result
}

// @lc code=end

func main() {
	root := BuildTree([]int{3, 9, 20, nil, nil, 15, 7})
	levelOrder(root)
}

func BuildTree(vals []int) (root *TreeNode) {
	root = &TreeNode{vals[0], nil, nil}
	for i := 1; i < len(vals); i++ {
		root.Left = &TreeNode{vals[i], nil, nil }
		root.Right = &TreeNode{Val: }
	}
}
