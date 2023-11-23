package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	var res int
	for _, val := range nums {
		res ^= val
	}
	return res
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
