package main

import (
	"fmt"
)

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// You can return the answer in any order.
//
// Example 1:
//
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:
//
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:
//
// Input: nums = [3,3], target = 6
// Output: [0,1]
//
// Constraints:
//
// 2 <= nums.length <= 104
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// Only one valid answer exists.
func main() {
	nums := []int{0, 3, 0, 2}
	target := 0
	fmt.Println(twoSum(nums, target))

}

func twoSum(array []int, target int) []int {
	seenNums := map[int]int{}
	for i, num := range array {
		potentialMatch := target - num
		if j, found := seenNums[potentialMatch]; found {
			return []int{j, i}
		}
		seenNums[num] = i
	}
	return []int{}
}
