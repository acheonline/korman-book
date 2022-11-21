package main

import "fmt"

//Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
//
//You must implement a solution with a linear runtime complexity and use only constant extra space.
//Constraints:
//
//1 <= nums.length <= 3 * 104
//-3 * 104 <= nums[i] <= 3 * 104
//Each element in the array appears twice except for one element which appears only once.

func main() {
	fmt.Println(singleNumber([]int{-1, -1, -1, -1, -1, -1, 1}))
}

func singleNumber(nums []int) int {
	seenNums := map[int]int{}
	var res int

	for _, num := range nums {
		seenNums[num] = seenNums[num] + 1
	}

	for k, v := range seenNums {
		if v == 1 {
			res = k
		}
	}
	return res
}
