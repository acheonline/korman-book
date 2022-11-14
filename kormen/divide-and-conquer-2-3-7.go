package main

import (
	"fmt"
	"kormen-book/libs"
)

func main() {

	array := libs.GenerateArray(4)           //O(n)
	fmt.Println("unsorted array is:", array) //O(0)
	libs.MergeSort(array, len(array))        // O(n*lg(n))
	fmt.Println("sorted array:", array)      //O(n)
	findSumInArray(array, array[2]+array[3])
}

func findSumInArray(arr []int, n int) { // 0(n)
	if len(arr) < 2 {
		return
	}

	mid := len(arr) / 2
	sum := arr[mid-1] + arr[mid]
	if sum == n {
		fmt.Println("got you:", arr[mid-1], "+", arr[mid], "=", n)
		return
	} else if sum > n {
		l := arr[:mid]
		findSumInArray(l, n)
	} else {
		r := arr[mid:]
		findSumInArray(r, n)
	}
}
