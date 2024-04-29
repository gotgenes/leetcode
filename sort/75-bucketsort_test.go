package sort_test

import (
	"testing"

	"github.com/gotgenes/leetcode/sort"
	"github.com/stretchr/testify/assert"
)

func TestBucketSort(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sort.BucketSort(nums)
	assert.Equal(t, []int{0, 0, 1, 1, 2, 2}, nums, "sorts integers")
}
