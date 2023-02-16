package main

import (
	"fmt"
)

var top = 0
var S = make([]int, 8)

func main() {
	fmt.Println(IsEmpty())
	Push(1)
	fmt.Println(IsEmpty())
	Push(2)
	fmt.Println(Pop())
	Push(3)
	Push(4)
	fmt.Println(Pop())
	fmt.Println(Pop())
	fmt.Println(Pop())
	fmt.Println(IsEmpty())
}

func IsEmpty() bool {
	is := top == 0
	return is
}

func Push(x int) {
	S[top] = x
	top++
}

func Pop() int {
	if IsEmpty() {
		panic("set is empty")
	} else {
		top--
		return S[top]
	}
}
