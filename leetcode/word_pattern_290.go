package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("actual:", wordPattern("abba", "dog cat cat dog"), ", expected: true")
	fmt.Println("actual:", wordPattern("abba", "dog cat cat fish"), ", expected: false")
	fmt.Println("actual:", wordPattern("aaaa", "dog cat cat dog"), ", expected: false")
	fmt.Println("actual:", wordPattern("abba", "dog dog dog dog"), ", expected: false")
}

func wordPattern(pattern string, s string) bool {
	ps := strings.Split(pattern, "")
	ss := strings.Split(s, " ")
	return isMatch(ps, ss) && isMatch(ss, ps)
}

func isMatch(s1, s2 []string) bool {
	size := len(s1)

	m := make(map[string]string, size)

	for i, r := range s1 {
		val, ok := m[r]
		if ok && !strings.EqualFold(s2[i], val) {
			return false
		} else {
			m[r] = s2[i]
		}
	}
	return true
}
