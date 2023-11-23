package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	memory := make(map[int]int) // val : position
	for pos, val := range nums {
		if m, ok := memory[val]; ok {
			check := m - pos
			if check < 0 {
				check *= -1
			}
			if check <= k {
				return true
			}
			memory[val] = pos
		} else {
			memory[val] = pos
		}
	}
	return false
}

func main() {
	nums := []int{1, 0, 1, 1} // 0 1 2 3
	k := 1
	fmt.Println(containsNearbyDuplicate(nums, k))
}
