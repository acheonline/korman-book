package main

import "fmt"

func main() {
	fmt.Println("actual:", hammingWeight(00000000000000000000000000001011), ",expected: 3")
	fmt.Println("actual:", hammingWeight(00000000000000000000000010000000), ",expected: 1")
	fmt.Println("actual:", hammingWeight(00000000000000000000000000000000), ",expected: 0")
}

func hammingWeight(num uint32) int {
	count := 0

	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			count++
		}

		num = num >> 1
	}

	return count
}
