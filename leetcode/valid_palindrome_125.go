package main

import (
	"fmt"
	"strings"
)

//A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
//
//Given a string s, return true if it is a palindrome, or false otherwise.
//Constraints:
//
//1 <= s.length <= 2 * 105
//s consists only of printable ASCII characters.

func main() {
	fmt.Println("Expected 'true', actual: ", isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println("Expected 'false', actual: ", isPalindrome("race a car"))
	fmt.Println("Expected 'true', actual: ", isPalindrome("."))
	fmt.Println("Expected 'false', actual: ", isPalindrome(""))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isChar(s[i]) {
			i++
		}
		for i < j && !isChar(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}
