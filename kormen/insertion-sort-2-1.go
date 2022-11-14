package main

import (
	"fmt"
	"kormen-book/libs"
)

func main() { //O(n2)
	x := libs.GenerateArray(11)       //O(n)
	fmt.Println("basic array is:", x) //O(0)
	var i, key int                    //O(0)
	for j := 0; j < len(x); j++ {     //O(n2)
		key = x[j]                             //O(0)
		for i = j - 1; i >= 0 && x[i] > key; { //O(n)
			//to make desc insertion sort you have to check 'x[i] < key' condition in loop
			x[i+1] = x[i] //O(0)
			i--           //O(0)
			x[i+1] = key  //O(0)
		}
	}
	fmt.Println("sorted array is:", x) //O(0)
}
