package testing

import "github.com/gotgenes/leetcode"

func NewLinkedListFromArray(values []int) *leetcode.ListNode {
	var head *leetcode.ListNode
	var node *leetcode.ListNode
	for _, val := range values {
		newNode := &leetcode.ListNode{Val: val}
		if node != nil {
			node.Next = newNode
			node = newNode
		} else {
			head = newNode
			node = newNode
		}
	}
	return head
}
