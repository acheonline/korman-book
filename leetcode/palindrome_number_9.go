package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Given an integer x, return true if x is palindrome integer.
//
// An integer is a palindrome when it reads the same backward as forward.
//
// For example, 121 is a palindrome while 123 is not.
// Constraints:
//
// -231 <= x <= 231 - 1

func main() {
	nums := []int{
		1, 2, 3, 4, 10, 11, 12, 13, 121, 123, -121, 0, 230, 232, 1000021}
	for _, n := range nums {
		fmt.Println("The number", n, "is palindrome:", checkIfNumberIsPalindrome(n))
	}
}

func checkIfNumberIsPalindrome(x int) bool {
	str := strings.Split(strconv.Itoa(x), "")
	if x >= 0 && x < 10 {
		return true
	}
	if x < 0 {
		return false
	}
	is := false
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] == str[length-1-i] {
			is = true
		} else {
			is = false
			break
		}
	}
	return is
}
