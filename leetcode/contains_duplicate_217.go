package main

import "fmt"

// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

//Constraints:
//
//1 <= nums.length <= 105
//-109 <= nums[i] <= 109

func main() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	isDuplicate := false
	m := make(map[int]int)
	for _, k := range nums {
		if m[k] > 0 {
			isDuplicate = true
			break
		}
		m[k] += 1
	}
	return isDuplicate
}
