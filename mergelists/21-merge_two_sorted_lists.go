package mergelists

import "github.com/gotgenes/leetcode"

func mergeTwoLists(list1 *leetcode.ListNode, list2 *leetcode.ListNode) *leetcode.ListNode {
	var head *leetcode.ListNode

	if list1 == nil {
		if list2 != nil {
			head = list2
			list2 = list2.Next
		}
	} else {
		if list2 == nil {
			head = list1
			list1 = list1.Next
		} else {
			if list1.Val < list2.Val {
				head = list1
				list1 = list1.Next
			} else {
				head = list2
				list2 = list2.Next
			}
		}
	}
	if head == nil {
		return nil
	}
	head.Next = nil

	current := head

	for list1 != nil {
		if list2 == nil {
			current.Next = list1
			list1 = list1.Next
		} else {
			if list1.Val < list2.Val {
				current.Next = list1
				list1 = list1.Next
			} else {
				current.Next = list2
				list2 = list2.Next
			}
		}
		current = current.Next
		current.Next = nil
	}
	if list2 != nil {
		current.Next = list2
	}
	return head
}
