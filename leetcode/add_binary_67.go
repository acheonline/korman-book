package main

import (
	"fmt"
)

//Given two binary strings a and b, return their sum as a binary string.
//Constraints:
//
//1 <= a.length, b.length <= 104
//a and b consist only of '0' or '1' characters.
//Each string does not contain leading zeros except for the zero itself.

func main() {
	var a, b = "1010", "1011"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		return addBinary(b, a)
	}
	difference := len(b) - len(a)

	// making both strings equal by adding 0 in front of a smaller string
	for i := 0; i < difference; i++ {
		a = "0" + a
	}
	// each addition
	carry := "0"
	// In this variable we will store our final string
	answer := ""

	// traversing the strings and adding them by picking the index from the end /*
	/* For example, we are adding “100” and ”110”.
	   So, for the last characters in the string i.e “0” and “0” the first else
	   if condition will run.
	      Then for the middle characters i.e “0” and “1” the last else if
	   condition will run and
	      for the first characters i.e “1” and “1” the first if condition will run.
	*/
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '1' && b[i] == '1' {
			if carry == "1" {
				answer = "1" + answer
			} else {
				answer = "0" + answer
				carry = "1"
			}
		} else if a[i] == '0' && b[i] == '0' {
			if carry == "1" {
				answer = "1" + answer
				carry = "0"
			} else {
				answer = "0" + answer
			}
		} else if a[i] != b[i] {
			if carry == "1" {
				answer = "0" + answer
			} else {
				answer = "1" + answer
			}
		}
	}
	if carry == "1" {
		answer = "1" + answer
	}
	return answer
}
