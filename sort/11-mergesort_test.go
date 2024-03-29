package sort_test

import (
	"testing"

	"github.com/gotgenes/leetcode/sort"
	"github.com/stretchr/testify/assert"
)

func TestHandlesEmptyLists(t *testing.T) {
	assert.Nil(t, sort.MergeSort(nil))
	assert.Empty(t, sort.MergeSort([]int{}))
}

func TestSortsElements(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4}, sort.MergeSort([]int{1, 2, 3, 4}), "already sorted stays sorted")
	assert.Equal(t, []int{1, 2, 3, 4}, sort.MergeSort([]int{4, 2, 3, 1}), "sorts unsorted")
}
