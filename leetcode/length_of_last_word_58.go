package main

import (
	"fmt"
	"strings"
)

// Given a string s consisting of words and spaces, return the length of the last word in the string.
// A word is a maximal substring consisting of non-space characters only.

// Constraints:
//
// 1 <= s.length <= 104
// s consists of only English letters and spaces ' '.
// There will be at least one word in s.
func main() {
	fmt.Println("expected 5, actual is: ", lengthOfLastWord("Hello World"))
	fmt.Println("expected 4, actual is: ", lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println("expected 6, actual is: ", lengthOfLastWord("luffy is still joyboy"))
	fmt.Println("expected 0, actual is: ", lengthOfLastWord(""))
	fmt.Println("expected 0, actual is: ", lengthOfLastWord("      "))
}
func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	arr := strings.Split(s, "")
	var found bool
	var count int
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == " " && !found {
			continue
		} else if arr[i] == " " && found {
			break
		} else if arr[i] != " " {
			found = true
			count += 1
		}
	}
	return count
}
