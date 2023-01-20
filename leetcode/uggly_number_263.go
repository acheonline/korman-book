package main

import "fmt"

//An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
//
//Given an integer n, return true if n is an ugly number.
//
//Constraints:
//
//-231 <= n <= 231 - 1

func main() {
	fmt.Println("actual:", isUgly(6), ", expected: true")
	fmt.Println("actual:", isUgly(1), ", expected: true")
	fmt.Println("actual:", isUgly(14), ", expected: false")
	fmt.Println("actual:", isUgly(8), ", expected: true")

}

func isUgly(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 || n == 2 || n == 3 || n == 5 {
		return true
	}
	if (n%2 == 0 && isUgly(n/2)) || (n%3 == 0 && isUgly(n/3)) || (n%5 == 0 && isUgly(n/5)) {
		return true
	}
	return false
}
