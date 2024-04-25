package kthlargest_test

import (
	"testing"

	"github.com/gotgenes/leetcode/kthlargest"
	"github.com/stretchr/testify/assert"
)

func TestKthLargest(t *testing.T) {
	result := kthlargest.FindKthLargest([]int{6, 2, 4, 1, 3, 5}, 3)
	assert.Equal(t, 4, result, "finds kth largest integer")
}
