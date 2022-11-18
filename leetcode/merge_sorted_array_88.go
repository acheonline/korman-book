package main

import "fmt"

var res []int

func main() {
	ints := merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	fmt.Println(ints)
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	if len(nums2) == 0 {
		res = nums1
		return nums1
	}

	if m == 0 {
		for k := 0; k < n; k++ {
			nums1[k] = nums2[k]
		}
		return nums1
	}

	k := len(nums1)
	for i := m - 1; i >= 0; i-- {
		nums1[k-1] = nums1[i]
		k--
	}

	i := n
	j := 0

	for k := 0; k < m+n; k++ {
		if i < n+m && j < n {
			if nums1[i] < nums2[j] {
				nums1[k] = nums1[i]
				i++
			} else {
				nums1[k] = nums2[j]
				j++
			}
		} else if j < n {
			nums1[k] = nums2[j]
			j++
		}
	}
	return nums1
}
