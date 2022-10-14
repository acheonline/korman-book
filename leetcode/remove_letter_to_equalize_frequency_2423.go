package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("correct string(true):", equalFrequency("acbc"))
	fmt.Println("correct string(false):", equalFrequency("aazz"))
	fmt.Println("2 chars(false):", equalFrequency("aa"))
	fmt.Println("1 char (false):", equalFrequency("a"))
}
func equalFrequency(word string) bool {
	length := len(word)

	for i := 0; i < length-1; i++ {
		if strings.Count(word, string(word[i])) == strings.Count(word, string(word[i+1])) && i < length-1 {
			continue
		} else {
			return true
		}
	}
	return false
}
