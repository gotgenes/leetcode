package sort

func MergeSort(nums []int) []int {
	numItems := len(nums)
	if numItems <= 1 {
		return nums
	}
	middle := numItems / 2
	left := MergeSort(nums[:middle])
	right := MergeSort(nums[middle:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	merged := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		l, r := left[i], right[j]
		if l <= r {
			merged = append(merged, l)
			i++
		} else {
			merged = append(merged, r)
			j++
		}
	}

	for ; i < len(left); i++ {
		merged = append(merged, left[i])
	}
	for ; j < len(right); j++ {
		merged = append(merged, right[j])
	}
	return merged
}
