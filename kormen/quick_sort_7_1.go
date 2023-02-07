package main

import (
	"fmt"
	"time"
)

func main() {
	a := []int{9, 16, 3, 12, 8, 4, 3, 2, 9, 5}
	a2 := []int{0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1}

	start := time.Now()
	ascQuickSort(a2, 0, len(a2)-1)
	elapsed := time.Since(start)
	fmt.Printf("Quick Sort Asc took %s\n", elapsed)
	fmt.Println(a2)

	start = time.Now()
	descQuickSort(a, 0, len(a)-1)
	elapsed = time.Since(start)
	fmt.Printf("Quick Sort Desc took %s\n", elapsed)
	fmt.Println(a)
}

// asc sort functions
func ascQuickSort(arr []int, low, high int) {
	if low < high {
		p := ascPartition(arr, low, high)
		ascQuickSort(arr, low, p-1)
		ascQuickSort(arr, p+1, high)
	}
}

func ascPartition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

// desc sort functions
func descQuickSort(arr []int, low, high int) {
	if high > low {
		p := descPartition(arr, low, high)
		descQuickSort(arr, low, p-1)
		descQuickSort(arr, p+1, high)
	}
}

func descPartition(arr []int, low, high int) int {
	l := arr[low]
	i := high
	for j := high; j > low; j-- {
		if arr[j] <= l {
			arr[j], arr[i] = arr[i], arr[j]
			i--
		}
	}
	arr[i], arr[low] = arr[low], arr[i]
	return i
}
