package p92

//https://leetcode.com/problems/reverse-linked-list-ii/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m >= n {
		return head
	}
	var prev *ListNode
	cur := head

	//move cur to the first one to be reversed
	for i := 1; i < m; i++ {
		prev = cur
		cur = cur.Next
	}

	startPoint := prev
	tail := cur

	for j := 0; j < (n-m)+1; j++ {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	if startPoint != nil {
		startPoint.Next = prev
	} else {
		head = prev
	}

	tail.Next = cur
	return head
}
