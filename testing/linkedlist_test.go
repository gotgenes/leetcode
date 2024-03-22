package testing_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gotgenes/leetcode"
	lt "github.com/gotgenes/leetcode/testing"
)

func TestCreatesEmptyList(t *testing.T) {
	node := lt.NewLinkedListFromArray([]int{})
	assert.Nil(t, node)
	node = lt.NewLinkedListFromArray(nil)
	assert.Nil(t, node)
}

func TestCreatesLinkedList(t *testing.T) {
	node := lt.NewLinkedListFromArray([]int{1, 2, 3})
	expected := &leetcode.ListNode{
		Val: 1,
		Next: &leetcode.ListNode{
			Val: 2,
			Next: &leetcode.ListNode{
				Val: 3,
			},
		},
	}
	assert.Equal(t, expected, node, "creates a linked list")
}
