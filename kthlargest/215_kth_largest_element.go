package kthlargest

func FindKthLargest(nums []int, k int) int {
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)
	return quickSelectKthLargest(numsCopy, 0, len(numsCopy)-1, k)
}

func quickSelectKthLargest(nums []int, startIndex int, endIndex int, k int) int {
	pivotValue := nums[endIndex]
	pivotIndex := startIndex
	for index := startIndex; index < endIndex; index++ {
		if nums[index] <= pivotValue {
			nums[pivotIndex], nums[index] = nums[index], nums[pivotIndex]
			pivotIndex++
		}
	}
	nums[endIndex], nums[pivotIndex] = nums[pivotIndex], nums[endIndex]

	kthPosition := len(nums) - k
	if kthPosition == pivotIndex {
		return pivotValue
	} else if kthPosition < pivotIndex {
		return quickSelectKthLargest(nums, startIndex, max(0, pivotIndex-1), k)
	} else {
		return quickSelectKthLargest(nums, pivotIndex, endIndex, k)
	}
}
