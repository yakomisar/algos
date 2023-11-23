package main

import "fmt"

func hammingWeight(num uint32) int {
	var sum int
	for num > 0 {
		if num%2 != 0 {
			sum++
		}
		num >>= 1
	}
	return sum
}

func main() {
	var n uint32
	n = 00000000000000000000000000001011
	fmt.Println(hammingWeight(n))
}
