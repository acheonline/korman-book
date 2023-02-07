package main

import (
	"fmt"
	"time"
)

func main() {
	a := []int{9, 16, 3, 12, 8, 4, 3, 2, 9, 5}
	a2 := []int{0, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1}

	start := time.Now()
	quickSort(a2, 0, len(a2)-1)
	elapsed := time.Since(start)
	fmt.Printf("Quick Sort 2 took %s\n", elapsed)
	fmt.Println(a2)

	start = time.Now()
	desQuickSort(a, 0, len(a)-1)
	elapsed = time.Since(start)
	fmt.Printf("Quick Sort 1 took %s\n", elapsed)
	fmt.Println(a)
}
func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
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

func desQuickSort(arr []int, low, high int) {
	if high > low {
		p := desPartition(arr, low, high)
		desQuickSort(arr, low, p-1)
		desQuickSort(arr, p+1, high)
	}
}

func desPartition(arr []int, low, high int) int {
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
