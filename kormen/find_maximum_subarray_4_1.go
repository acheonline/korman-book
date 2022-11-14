package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []int{1, 2, 1, 2, 9}
	maxSum := maxSubArraySum(&arr, 0, len(arr)-1)
	fmt.Println("Maximum contiguous sum is", maxSum)
}

// Returns sum of maximum sum subarray in aa[l..h]
func maxSubArraySum(arr *[]int, l, h int) int {
	//Invalid Range: low is greater than high
	if l > h {
		return math.MinInt
	}
	// Base Case: Only one element
	if l == h {
		return (*arr)[l]
	}

	// Find middle point
	m := (l + h) / 2

	/* Return maximum of following three
	   possible cases:
	   a) Maximum subarray sum in left half
	   b) Maximum subarray sum in right half
	   c) Maximum subarray sum such that the
	   subarray crosses the midpoint */
	ml := maxSubArraySum(arr, l, m-1)
	mr := maxSubArraySum(arr, m+1, h)
	return max(max(ml, mr), maxCrossingSum(arr, l, m, h))
}

// Find the maximum possible sum in arr[]
// such that arr[m] is part of it
func maxCrossingSum(arr *[]int, l, m, h int) int {
	// Include elements on left of mid.
	sum := 0
	leftSum := math.MinInt
	for i := m; i >= l; i-- {
		sum = sum + (*arr)[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	// Include elements on right of mid
	sum = 0
	rightSum := math.MinInt
	for i := m; i <= h; i++ {
		sum = sum + (*arr)[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	// Return sum of elements on left
	// and right of mid
	// returning only left_sum + right_sum will fail for
	// [-2, 1]
	return max(leftSum+rightSum-(*arr)[m],
		max(leftSum, rightSum))
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
