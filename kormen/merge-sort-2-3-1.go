package main

import (
	"fmt"
	"kormen-book/libs"
)

func main() {

	array := libs.GenerateArray(9)           //O(n)
	fmt.Println("unsorted array is:", array) //O(0)
	mergeSort(array, len(array))             // 2*O(n/2) + O(n)
	fmt.Println("sorted array:", array)
}

func mergeSort(arr []int, n int) { // O(n)
	if n < 2 { // if size less than 2 there is no need to make mergeSort
		return
	}

	mid := n / 2
	l := make([]int, mid)
	r := make([]int, n-mid)
	for i := 0; i < mid; i++ { // O(n)
		l[i] = arr[i]
	}
	for i := 0; i < n-mid; i++ { // O(n)
		r[i] = arr[mid+i]
	}

	mergeSort(l, mid)   // O(n/2)
	mergeSort(r, n-mid) // O(n/2)

	merge(arr, l, r, mid, n-mid) // O(n)
}

func merge(arMain, l, r []int, left, right int) { // O(n)

	if left < 0 || right < 0 {
		panic("array bounds are negative")
	}

	var i, j, k = 0, 0, 0
	for i < left && j < right { // O(n)
		if l[i] <= r[j] {
			arMain[k] = l[i]
			k++
			i++
		} else {
			arMain[k] = r[j]
			k++
			j++
		}
	}

	for i < left { // O(n)
		arMain[k] = l[i]
		k++
		i++
	}

	for j < right { // O(n)
		arMain[k] = r[j]
		k++
		j++
	}
}
