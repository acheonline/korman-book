package main

import (
	"fmt"
)

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
