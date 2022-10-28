package main

import "fmt"

//Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
//You must write an algorithm with O(log n) runtime complexity.

var count int

func main() {

	arr := []int{1, 3, 5, 6}
	//	target := 2
	for i := 0; i < 10; i++ {
		count = 0
		fmt.Println(i, searchInsert(arr, i))
	}
}

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	var mid int

	for low <= high {
		mid = (low + high + 1) / 2
		midVal := nums[mid]

		if midVal < target {
			low = mid + 1
		} else if midVal > target {
			high = mid - 1
		} else {
			return mid
		} // key found
	}
	return low // key not found.
}
