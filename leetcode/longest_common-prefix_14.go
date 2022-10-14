package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := []string{
		"flower", "flow", "flight",
	}

	s2 := []string{
		"dog", "racecar", "car",
	}

	fmt.Println(longestCommonPrefix(s1))
	fmt.Println(longestCommonPrefix(s2))

}
func longestCommonPrefix(strs []string) string {
	var longestPrefix = ""
	var endPrefix = false

	if len(strs) > 0 {
		sort.Strings(strs)
		first := strs[0]
		last := strs[len(strs)-1]
		for i := 0; i < len(first); i++ {
			if !endPrefix && last[i] == first[i] {
				longestPrefix = longestPrefix + string(last[i])
			} else {
				endPrefix = true
			}
		}
	}
	return longestPrefix
}
