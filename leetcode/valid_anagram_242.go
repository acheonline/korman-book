package main

import (
	"fmt"
	"sort"
)

//Given two strings s and t, return true if t is an anagram of s, and false otherwise.
//
//An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
//Constraints:
//
//1 <= s.length, t.length <= 5 * 104
//s and t consist of lowercase English letters.

func main() {
	fmt.Println("actual", isAnagram("anagram", "nagaram"), "/ expected true")
	fmt.Println("actual", isAnagram("rat", "car"), "/ expected false")
	fmt.Println("actual", isAnagram("aa", "bb"), "/ expected false")
	fmt.Println("actual", isAnagram("a", "ab"), "/ expected false")

}

func isAnagram(s string, t string) bool {
	is := true
	if len(s) != len(t) {
		return false
	}
	r1 := []rune(s)
	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})
	r2 := []rune(t)
	sort.Slice(r2, func(i, j int) bool {
		return r2[i] < r2[j]
	})
	for i, r := range r1 {
		if r != r2[i] {
			is = false
			break
		}
	}
	return is
}
