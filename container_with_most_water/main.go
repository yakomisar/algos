package main

import "fmt"

func getMin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func maxArea(height []int) int {
	totalSum := 0
	for i, j := 0, len(height)-1; i < j; {
		min := getMin(height[i], height[j])
		res := min * (j - i)
		if totalSum <= res {
			totalSum = res
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return totalSum
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
