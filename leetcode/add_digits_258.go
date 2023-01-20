package main

import "fmt"

// Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.
// Constraints:
//
// 0 <= num <= 231 - 1
func main() {
	fmt.Println("actual:", addDigits(38), ", expected: 2")
	fmt.Println("actual:", addDigits(0), ", expected: 0")
	fmt.Println("actual:", addDigits(9999999), ", expected: 9")

}

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	for num >= 10 {
		n := num
		num = 0
		for n > 0 {
			num += n % 10
			n = n / 10
		}
	}
	return num
}
