package main

import "fmt"

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	tmp := make(map[int]int)
	var result int
	var count int
	for _, val := range nums {
		if _, ok := tmp[val]; ok {
			tmp[val]++
		} else {
			tmp[val] = 1
		}
	}
	for k, v := range tmp {
		if v > count {
			result = k
			count = v
		}
	}
	return result
}

func main() {
	test1 := []int{8, 8, 7, 7, 7}
	fmt.Println(majorityElement(test1))
}
