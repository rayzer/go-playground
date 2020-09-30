package main

import "fmt"

/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */
//BFS version
//0 ms, faster than 100.00%
// @lc code=start

type node struct {
	s           string
	left, right int
}

func generateParenthesisBFS(n int) []string {
	var result []string
	queue := []node{}
	queue = append(queue, node{"(", 1, 0})
	for len(queue) != 0 {
		fmt.Println(queue)
		size := len(queue)
		for _, e := range queue {
			if e.left < n {
				queue = append(queue, node{e.s + "(", e.left + 1, e.right})
			}
			if e.right < e.left {
				queue = append(queue, node{e.s + ")", e.left, e.right + 1})
			}
			if e.left == n && e.right == n {
				result = append(result, e.s)
			}
		}
		queue = queue[size:]
	}
	return result
}

// @lc code=end

func main() {
	fmt.Println(generateParenthesisBFS(3))
}
