package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(25))
	fmt.Println(convertToTitle(26))
	fmt.Println(convertToTitle(28))

	fmt.Println(convertToTitle(701))
	fmt.Println(convertToTitle(52))
}

func convertToTitle(columnNumber int) string {
	res := ""
	for columnNumber > 0 {
		rem := columnNumber % 26
		if rem == 0 {
			res = "Z" + res
			columnNumber = (columnNumber / 26) - 1
		} else {
			res = string(rune(rem-1)+'A') + res
			columnNumber = columnNumber / 26
		}
	}
	return res
}
