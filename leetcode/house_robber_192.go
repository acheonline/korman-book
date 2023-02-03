package main

import "fmt"

func main() {
	fmt.Println("actual:", rob([]int{1, 2, 3, 1}), ",expected: 4")
	fmt.Println("actual:", rob([]int{2, 7, 9, 3, 1}), ",expected: 12")
	fmt.Println("actual:", rob([]int{2, 1, 1, 2}), ",expected: 4")

}

func rob(nums []int) int {
	lenNums := len(nums)

	if lenNums == 0 {
		return 0
	}

	maxMoney := make([]int, lenNums)

	maxMoney[0] = nums[0]

	if lenNums > 1 {
		maxMoney[1] = nums[1]
	}

	if lenNums > 2 {
		maxMoney[2] = nums[2] + nums[0]
	}

	for i := 3; i < lenNums; i++ {
		if maxMoney[i-2] > maxMoney[i-3] {
			maxMoney[i] = nums[i] + maxMoney[i-2]
		} else {
			maxMoney[i] = nums[i] + maxMoney[i-3]
		}

	}

	max := 0
	for i := lenNums - 1; i >= 0; i-- {
		if maxMoney[i] > max {
			max = maxMoney[i]
		}
	}

	return max
}
