package main

import "fmt"

//Given an array nums of size n, return the majority element.
//
//The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

//Constraints:
//
//n == nums.length
//1 <= n <= 5 * 104
//-109 <= nums[i] <= 109

//Follow-up: Could you solve the problem in linear time and in O(1) space?

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majorityElement([]int{6, 5, 5}))

}

func majorityElement(nums []int) int {
	count := 0
	candidate := -1

	// Finding majority candidate
	for _, elem := range nums {
		if count == 0 {
			candidate = elem
			count = 1
		} else {
			if elem == candidate {
				count++
			} else {
				count--
			}
		}
	}

	// Checking if majority candidate occurs more than
	// n/2 times
	count = 0
	for _, elem := range nums {
		if elem == candidate {
			count++
		}
	}
	if count > (len(nums) / 2) {
		return candidate
	}
	return -1

	// The last for loop and the if statement step can
	// be skip if a majority element is confirmed to
	// be present in an array just return candidate
	// in that case
	//candidate := nums[0]
	//counter := 1
	//for _, elem := range nums {
	//	if counter == 0 {
	//		counter = 1
	//		candidate = elem
	//	} else {
	//		if candidate == elem {
	//			counter++
	//		} else {
	//			counter--
	//		}
	//	}
	//}
	//return candidate
}
