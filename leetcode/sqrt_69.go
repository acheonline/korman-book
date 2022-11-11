package main

import (
	"fmt"
	"math"
)

//Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.
//
//You must not use any built-in exponent function or operator.
//
//For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.

//Constraints:
//
//0 <= x <= 231 - 1

func main() {
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	z := 100.0

	step := func() float64 {
		return z - (z*z-float64(x))/(2*z)
	}

	for zz := step(); math.Abs(zz-z) > 0.0000001; {
		z = zz
		zz = step()
	}
	return int(z)
}
