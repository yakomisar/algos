package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for i := range nums {
		left := target - nums[i]
		if val, ok := hashTable[left]; ok {
			return []int{val, i}
		} else {
			hashTable[nums[i]] = i
		}
	}
	return []int{}
}

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(arr, target))
}
