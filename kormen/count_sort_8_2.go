package main

import "fmt"

func main() {
	fmt.Println(countSort([]int{9, 16, 3, 12, 8, 4, 1, 2, 8, 5}, 17))
}

func countSort(a []int, k int) []int {
	l := len(a)
	b := make([]int, l)
	c := make([]int, k+1)
	for j := 0; j < l; j++ {
		c[a[j]]++
	}
	//Now Ci contains number of elements equal i

	for i := 1; i < k; i++ {
		c[i] += c[i-1]
	}
	//Now Ci contains number of elements less or equal then i
	for j := l - 1; j >= 0; j-- {
		b[c[a[j]]-1] = a[j]
	}
	return b
}
