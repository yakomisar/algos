package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return []int{}
	}
	left := 0
	right := len(numbers) - 1
	for left < right {
		res := numbers[left] + numbers[right]
		if res == target {
			return []int{left + 1, right + 1}
		} else if res > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(numbers, target))
}
