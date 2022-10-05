package main

import (
	"fmt"
	"kormen-book/libs"
)

func main() {

	array := libs.GenerateArray(9)           //O(n)
	fmt.Println("unsorted array is:", array) //O(0)
	mergeSortEnd(array, len(array))          // 2*O(n/2) + O(n)
	fmt.Println("sorted array:", array)
}

func mergeSortEnd(arr []int, n int) { // O(n)
	if n < 2 {
		return
	}

	pos := n - 1
	l := make([]int, pos)
	r := make([]int, 1)
	for i := 0; i < pos; i++ { // O(n)
		l[i] = arr[i]
	}
	r[0] = arr[pos]

	mergeSortEnd(l, pos)   // O(n/2)
	mergeSortEnd(r, n-pos) // O(n/2)

	mergeModified(arr, l, r, pos, n-pos) // O(n)
}

func mergeModified(arMain, l, r []int, left, right int) { // O(n)

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
