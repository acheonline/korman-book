package main

import "fmt"

func main() {
	fmt.Println(commonFactors(25, 30))
}

func commonFactors(a int, b int) int {
	var min, count int
	if a < b {
		min = a
	} else {
		min = b
	}
	for i := 1; i <= min; i++ {
		if a%i == 0 && b%i == 0 {
			count++
		}
	}

	return count
}
