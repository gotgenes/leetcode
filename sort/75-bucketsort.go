package sort

func BucketSort(nums []int) {
	capacity := -1
	var counts []int
	for _, value := range nums {
		if value > capacity {
			extensionLen := value - capacity + 1
			extension := make([]int, extensionLen)
			counts = append(counts, extension...)
			capacity = value
		}
		counts[value]++
	}
	i := 0
	for value, count := range counts {
		for j := 0; j < count; j++ {
			nums[i] = value
			i++
		}
	}
}
