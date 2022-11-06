package main

import "fmt"

//You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer.
//The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.
//Increment the large integer by one and return the resulting array of digits.

//Constraints:
//
//1 <= digits.length <= 100
//0 <= digits[i] <= 9
//digits does not contain any leading 0's.

func main() {

	arr1 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("expected [9,8,7,6,5,4,3,2,1,1], actual:", plusOne(arr1))

}
func plusOne(digits []int) []int {
	length := len(digits)
	res := make([]int, length)

	if length == 0 {
		panic("array is empty")
	}

	dozen, last := false, true
	for i := length - 1; i >= 0; i-- {
		if digits[i] == 9 && !dozen && !last {
			res[i] = digits[i]
		} else if dozen && digits[i] != 9 {
			res[i] = digits[i] + 1
			dozen = false
		} else if dozen && digits[i] == 9 {
			res[i] = 0
		} else if digits[i] != 9 && !dozen && !last {
			res[i] = digits[i]
		} else if digits[i] == 9 && !dozen && last {
			res[i] = 0
			dozen = true
			last = false
		} else if digits[i] == 9 && last {
			res[i] = digits[i]
		} else if !dozen && last {
			res[i] = digits[i] + 1
			last = false
		}
	}
	if digits[0] == 9 && dozen {
		res = append([]int{1}, res...)
	}
	return res
}
