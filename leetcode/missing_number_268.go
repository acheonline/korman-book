package main

import "fmt"

// Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.
// Constraints:
//
// n == nums.length
// 1 <= n <= 104
// 0 <= nums[i] <= n
// All the numbers of nums are unique.
func main() {
	fmt.Println("actual:", missingNumber([]int{3, 0, 1}), ", expected: 2")
	fmt.Println("actual:", missingNumber([]int{0, 1}), ", expected: 2")
	fmt.Println("actual:", missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}), ", expected: 8")
	fmt.Println("actual:", missingNumber([]int{1}), ", expected: 0")

}
func missingNumber(nums []int) int {
	sorted := make([]int, len(nums)+1)
	for _, num := range nums {
		sorted[num] = num
	}
	for n := len(sorted) - 1; n > 0; n-- {
		if sorted[n] == 0 {
			return n
		}
	}
	return 0
}
