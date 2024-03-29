package sort_test

import (
	"testing"

	"github.com/gotgenes/leetcode/sort"
	"github.com/stretchr/testify/assert"
)

func TestQuickSortHandlesEmptyLists(t *testing.T) {
	var vals []int
	sort.QuickSort(vals, 0, 0)
	assert.Nil(t, vals)
	vals = []int{}
	sort.QuickSort(vals, 0, 0)
	assert.Empty(t, vals)
}

func TestQuickSortSortsElements(t *testing.T) {
	vals := []int{1, 2, 3, 4}
	sort.QuickSort(vals, 0, 3)
	assert.Equal(t, []int{1, 2, 3, 4}, vals, "already sorted stays sorted")
	vals = []int{4, 2, 3, 1}
	sort.QuickSort(vals, 0, 3)
	assert.Equal(t, []int{1, 2, 3, 4}, vals, "sorts unsorted")
}
