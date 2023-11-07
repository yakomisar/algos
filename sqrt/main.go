package main

import (
	"fmt"
)

func mySqrt(x int) int {
	res := 0
	if x < 1 {
		return res
	}
	for res < x/2 {
		res = res * res
		if res == x {
			return res
		} else if (res+1)*(res+1) > x {
			return res
		}
		res += 1
	}
	return res
}

func main() {
	// x := 4
	// fmt.Println(mySqrt(x))
	// x = 8
	// fmt.Println(mySqrt(x))
	x := 9
	fmt.Println(mySqrt(x))
}
