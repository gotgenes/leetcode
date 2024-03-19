package sort_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/gotgenes/leetcode/sort"
)

var _ = Describe("MergeSort", func() {
	It("should handle empty lists", func() {
		Expect(sort.MergeSort(nil)).To(BeNil())
		Expect(sort.MergeSort([]int{})).To(Equal([]int{}))
	})
	It("should sort elements", func() {
		Expect(sort.MergeSort([]int{1, 2, 3, 4})).To(Equal([]int{1, 2, 3, 4}))
		Expect(sort.MergeSort([]int{4, 2, 3, 1})).To(Equal([]int{1, 2, 3, 4}))
	})
})
