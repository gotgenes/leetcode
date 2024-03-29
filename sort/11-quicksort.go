package sort

func QuickSort(vals []int, start int, end int) {
	if vals == nil || len(vals) <= 1 || (end-start) <= 1 {
		return
	}
	pivot := vals[end]
	left := start
	for i := start; i < end; i++ {
		if vals[i] < pivot {
			vals[left], vals[i] = vals[i], vals[left]
			left++
		}
	}
	vals[left], vals[end] = vals[end], vals[left]
	QuickSort(vals, start, left-1)
	QuickSort(vals, left+1, end)
}
