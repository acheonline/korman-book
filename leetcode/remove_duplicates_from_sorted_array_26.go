package main

import (
	"fmt"
	"strconv"
)

//Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same.
//Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.
//Return k after placing the final result in the first k slots of nums.
//Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}                                // Input array
	expectedNums := []string{"0", "1", "2", "3", "4", "_", "_", "_", "_", "_"} // The expected answer with correct length

	k := removeDuplicates(nums) // Calls your implementation

	if k == len(expectedNums) {
		panic("arrays' lengths are not equal")

	}
	for i := 0; i < k; i++ {
		if strconv.Itoa(nums[i]) != expectedNums[i] {
			fmt.Println(i)
			panic("arrays' items are not equal")
		}
	}
	fmt.Println(k)
}

func removeDuplicates(nums []int) int {
	if len(nums) > 0 {
		length, newLength := len(nums), 1
		for i := 1; i < length; i++ {
			if nums[i-1] != nums[i] {
				nums[newLength] = nums[i]
				newLength++
			}
		}
		return newLength
	}
	return 0
}
