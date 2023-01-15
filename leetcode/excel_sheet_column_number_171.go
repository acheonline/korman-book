package main

import (
	"fmt"
	"math"
)

// Given a string columnTitle that represents the column title as appears in an Excel sheet, return its corresponding column number.

//Constraints:

//1 <= columnTitle.length <= 7
//columnTitle consists only of uppercase English letters.
//columnTitle is in the range ["A", "FXSHRXW"].

func main() {
	fmt.Println("actual", titleToNumber("A"), "/ expected 1")
	fmt.Println("actual", titleToNumber("AB"), "/ expected 28")
	fmt.Println("actual", titleToNumber("ZY"), "/ expected 701")

}

func titleToNumber(columnTitle string) int {
	st := []rune(columnTitle)
	res := 0
	for index := len(st) - 1; index >= 0; index-- {
		pow := math.Pow(26, float64(len(st)-1-index))
		res += int(st[index]) % 64 * int(pow)
	}
	return res
}
