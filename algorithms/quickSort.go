package main

import (
	"fmt"
)

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func partition(nums *[]int, start, end int) int {
	pivot := start
	minThan := start + 1
	for i := start + 1; i <= end; i++ {
		if (*nums)[pivot] >= (*nums)[i] {
			swap(&(*nums)[minThan], &(*nums)[i])
			minThan++
		}
	}
	swap(&(*nums)[pivot], &(*nums)[minThan-1])
	pivot = minThan - 1
	return pivot
}

func QuickSort(nums *[]int, start, end int) {
	if start >= end {
		return
	}
	mid := partition(nums, start, end)
	QuickSort(nums, start, mid-1)
	QuickSort(nums, mid+1, end)
}

func main() {
	nums := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	QuickSort(&nums, 0, len(nums)-1)
	fmt.Println(nums)
}
