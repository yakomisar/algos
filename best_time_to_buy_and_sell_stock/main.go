package main

import "fmt"

func maxProfit(prices []int) int {
	var result int
	l := 0
	r := 1
	if len(prices) == 1 {
		return result
	} else {
		for r < len(prices) {
			tmp := prices[r] - prices[l]
			if tmp > result {
				result = tmp
				r++
				continue
			} else if prices[r] < prices[l] {
				l = r
			}
			r++
		}
	}
	return result
}

func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	prices2 := []int{7, 6, 4, 3, 1}
	prices3 := []int{7, 0, 5, 1}
	array := [][]int{prices1, prices2, prices3}
	for i := range array {
		fmt.Println(maxProfit(array[i]))
	}
}
