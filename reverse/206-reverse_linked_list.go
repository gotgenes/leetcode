package reverse

import "github.com/gotgenes/leetcode"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *leetcode.ListNode) *leetcode.ListNode {
	var prev *leetcode.ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
