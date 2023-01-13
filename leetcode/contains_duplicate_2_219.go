package main

import (
	"fmt"
)

//Given an integer array nums and an integer k,
//return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

//Constraints:
//1 <= nums.length <= 105
//-109 <= nums[i] <= 109
//0 <= k <= 105

func main() {
	nums := []int{1, 0, 1, 1}
	fmt.Println(containsDuplicate2(nums, 1))
}

func containsDuplicate2(nums []int, k int) bool {
	isDuplicate := false
	m := make(map[int]int)
	for i, key := range nums {
		val, ok := m[key]
		if ok && Abs(i-val) <= k {
			isDuplicate = true
		}
		m[key] = i
	}
	return isDuplicate
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
