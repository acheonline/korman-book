package main

import (
	"fmt"
)

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
// constrains:
// s consists of parentheses only '()[]{}'.
func main() {
	array := []string{
		"(", "()", "[]{}", "()[]{}", "(]",
		"([)]",
		"", "123", "__",
	}
	for _, s := range array {
		fmt.Println("The string:", s, "validity is:", isValid(s))
	}
}

//  1) not in pattern regexp
// 2) Open brackets must be closed by the same type of brackets - equality of number of brackets same types
// 3) Open brackets must be closed in the correct order - index of open brackets is less than index of closed brackets

func isValid(s string) bool {
	// if the string isn't of even length,
	// it can't be valid so return
	if len(s)%2 != 0 {
		return false
	}

	// set up stack and map
	var st []rune
	open := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	// loop over string
	for _, r := range s {

		// if the current character is in the open map,
		// put its closer into the stack and continue
		if closer, ok := open[r]; ok {
			st = append(st, closer)
			continue
		}

		// otherwise, we're dealing with a closer
		// check to make sure the stack isn't empty
		// and whether the top of the stack matches
		// the current character
		l := len(st) - 1
		if l < 0 || r != st[l] {
			return false
		}

		// take the last element off the stack
		st = st[:l]
	}

	// if the stack is empty, return true, otherwise false
	return len(st) == 0
}
