package mergelists

import "github.com/gotgenes/leetcode"

func MergeKlists(lists []*leetcode.ListNode) *leetcode.ListNode {
	if len(lists) == 0 {
		return &leetcode.ListNode{}
	}
	var head *leetcode.ListNode
	var tail *leetcode.ListNode
	numLists := len(lists)
	numEmpty := 0
	for numEmpty < numLists {
		numEmpty = 0
		var minValue int
		listIndex := numLists
		for i, node := range lists {
			if node == nil {
				numEmpty++
			} else {
				if (listIndex == numLists) || (node.Val < minValue) {
					minValue = node.Val
					listIndex = i
				}
			}
		}
		if numEmpty != numLists {
			nextNode := &leetcode.ListNode{
				Val: minValue,
			}
			if tail == nil {
				head = nextNode
				tail = nextNode
			} else {
				tail.Next = nextNode
				tail = nextNode
			}
			lists[listIndex] = lists[listIndex].Next
		}
	}
	return head
}
