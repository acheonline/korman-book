package libs

import (
	"math/rand"
	"time"
)

func GenerateArray(size int) []int {
	rand.Seed(time.Now().UnixNano()) //O(0)
	if size <= 0 {
		panic("Size can't be negative Int or equal Zero")
	}
	arr := make([]int, size)        //O(n)
	for i := 0; i < len(arr); i++ { //O(n)
		arr[i] = rand.Intn(100) //O(0)
	}
	return arr //O(0)
}
