package main

/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
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

//1：HashSet
func hasCycleSetVersion(head *ListNode) bool {
	nodePassed := make(map[*ListNode]bool)
	nodePassed[head] = true
	var ok bool
	for head != nil {
		head = head.Next
		_, ok = nodePassed[head]
		if !ok {
			nodePassed[head] = true
		} else {
			return true
		}
	}
	return false
}

//2: 快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// @lc code=end
