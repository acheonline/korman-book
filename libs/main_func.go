package libs

import (
	"math/rand"
	"time"
)

func GenerateArray(size int) []int {
	rand.Seed(time.Now().UnixNano()) //O(0)
	if size <= 0 {
		panic("Size can't be negative Int or equal Zero")
	}
	arr := make([]int, size)        //O(n)
	for i := 0; i < len(arr); i++ { //O(n)
		arr[i] = rand.Intn(100) //O(0)
	}
	return arr //O(0)
}

func MergeSort(arr []int, n int) { // O(n)
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

	MergeSort(l, mid)   // O(n/2)
	MergeSort(r, n-mid) // O(n/2)

	Merge(arr, l, r, mid, n-mid) // O(n)
}

func Merge(arMain, l, r []int, left, right int) { // O(n)

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
