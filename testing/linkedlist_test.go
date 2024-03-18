package testing_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/gotgenes/leetcode"
	"github.com/gotgenes/leetcode/testing"
)

var _ = Describe("linkedlist helpers", func() {
	It("can create an empty linked list", func() {
		node := testing.NewLinkedListFromArray([]int{})
		Expect(node).To(BeNil())
		node = testing.NewLinkedListFromArray(nil)
		Expect(node).To(BeNil())
	})

	It("can create a linked list", func() {
		node := testing.NewLinkedListFromArray([]int{1, 2, 3})
		expected := &leetcode.ListNode{
			Val: 1,
			Next: &leetcode.ListNode{
				Val: 2,
				Next: &leetcode.ListNode{
					Val: 3,
				},
			},
		}
		Expect(node).To(Equal(expected))
	})
})
