package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=547 lang=golang
 *
 * [547] Friend Circles
 */

// @lc code=start
func findCircleNum4(M [][]int) int {
	n := len(M)
	count := n

	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	find := func(a int) int {
		for a != parent[a] {
			parent[a] = parent[parent[a]] //压缩
			a = parent[a]                 //不断向上
		}
		return a
	}

	union := func(a, b int) {
		rootA := find(a)
		rootB := find(b)
		if rootA == rootB {
			return
		}
		parent[rootA] = rootB
		count--
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if M[i][j] == 1 {
				union(i, j)
			}
		}
	}
	return count
}

func findCircleNum(M [][]int) int {
	n := len(M)
	count := len(M)
	//init
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	find := func(p int) int {
		for p != parent[p] {
			parent[p] = parent[parent[p]] //压缩
			p = parent[p]
		}
		return p
	}

	union := func(p, q int) {
		rootP := find(p)
		rootQ := find(q)
		if rootP == rootQ {
			return
		}
		parent[rootP] = rootQ
		count--
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if M[i][j] == 1 {
				union(i, j)
			}
		}
	}
	return count
}

// @lc code=end

func main() {
	input := [][]int{
		{1, 1, 0}, {1, 1, 0}, {0, 0, 1},
	}
	fmt.Println(findCircleNum(input)) //2
}
