package main

import (
	"fmt"
)

func main() {
	fmt.Println(rle("ABBBBCD"))
}

func rle(s string) string {
	rs := []rune(s)
	l := len(rs) - 2
	res := ""
	cnt := 1
	for i := 0; i <= l; i++ {
		if rs[i] == rs[i+1] {
			if i != l {
				cnt++
			} else {
				res += string(rs[i]) + string(rune(49+cnt))
			}
		} else {
			res += string(rs[i]) + string(rune(48+cnt))
			if i != l {
				cnt = 1
			} else {
				res += string(rs[i+1]) + string(rune(49))
			}
		}
	}
	return res
}
