package main

import (
	"fmt"
	"kormen-book/libs"
	"time"
)

func main() {
	array := libs.GenerateArray(11) //O(n)
	for i := 0; i < len(array); i++ {
		ds := time.Now()
		sort(array, 0, len(array)-1, i)
		de := time.Now()
		elapsed := de.Sub(ds)
		fmt.Println("Threshold:", i, ". Time spent:", elapsed.Nanoseconds(), ",ns")
	}
}

func sort(arr []int, p, r, threshold int) {
	if r-p > threshold {
		var n = len(arr)
		if n < 2 { // if size less than 2 there is no need to make mergeSort
			return
		}

		mid := n / 2
		al := make([]int, mid)
		ar := make([]int, n-mid)
		for i := 0; i < mid; i++ { // O(n)
			al[i] = arr[i]
		}
		for i := 0; i < n-mid; i++ { // O(n)
			ar[i] = arr[mid+i]
		}

		sort(al, p, mid, threshold)   // O(n/2)
		sort(ar, mid+1, r, threshold) // O(n/2)

		libs.Merge(arr, al, ar, mid, n-mid)
	} else {
		libs.InsertionSort(arr)
	}
}
