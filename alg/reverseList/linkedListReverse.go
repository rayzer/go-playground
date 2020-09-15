package main

import (
	"fmt"
)

type node struct {
	data string
	next *node
}

func main() {
	var head = &node{"a", nil}
	head.next = &node{"b", nil}
	head.next.next = &node{"c", nil}
	head.next.next.next = &node{"d", nil}
	head.next.next.next.next = &node{"e", nil}

	printList(head)
	println()
	reversed := reverse(head)
	printList(reversed)
	println()
	reversed = rReverse(reversed)
	printList(reversed)
}

//递归函数：
func rReverse(head *node) *node {
	if head.next == nil {
		return head
	}
	last := rReverse(head.next)
	head.next.next = head //反转
	head.next = nil       //尾巴
	return last
}

func reverse(n *node) *node {
	var prev *node
	for n != nil {
		t := n.next
		n.next = prev
		prev, n = n, t
	}
	return prev
}

func printList(head *node) {
	current := head
	for current != nil {
		fmt.Print(current.data)
		if current.next != nil {
			print("->")
		}
		current = current.next
	}
}
