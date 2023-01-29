package main

import (
	"fmt"
)

func main() {
	//fmt.Println("actual:", isHappy(19), ", expected: true")
	//fmt.Println("actual:", isHappy(2), ", expected: false")
	//fmt.Println("actual:", isHappy(1111111), ",expected: true")
	fmt.Println("actual:", isHappy(4), ",expected: true")

}

func isHappy(n int) bool {
	if n == 2 || n == 3 || n == 4 {
		return false
	}
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	if sum == 1 {
		return true
	} else {
		return isHappy(sum)
	}
}
