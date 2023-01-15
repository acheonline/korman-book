package main

import "fmt"

//Given an integer n, return true if it is a power of two. Otherwise, return false.
//
//An integer n is a power of two, if there exists an integer x such that n == 2x.

func main() {
	fmt.Println("actual", isPowerOfTwo(1), "/ expected true")
	fmt.Println("actual", isPowerOfTwo(16), "/ expected true")
	fmt.Println("actual", isPowerOfTwo(3), "/ expected false")
}

func isPowerOfTwo(n int) bool {
	pot := true
	if n == 1 {
		return true
	} else if n <= 0 {
		return false
	} else if n > 0 {
		for pot {
			if n%2 > 0 && n >= 2 {
				pot = false
			} else if n == 2 {
				break
			} else {
				n = n / 2
			}
		}
	}
	return pot
}
