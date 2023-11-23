package main

import "fmt"

func climbStairs(n int) int {
	steps := 0
	if n < 1 {
		return steps
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	st := 1
	end := 1
	for i := 2; i <= n; n-- {
		steps = st + end
		end = st
		st = steps
	}
	return steps
}

func main() {
	n := 3
	fmt.Println(climbStairs(n))
}
