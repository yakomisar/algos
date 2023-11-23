package main

import "fmt"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	memory := make([]int, 0, len(nums))
	memory = append(memory, nums[0])
	memory = append(memory, max(nums[0], nums[1]))
	for i := 2; i < len(nums); i++ {
		memory = append(memory, max(memory[i-1], memory[i-2]+nums[i]))
	}
	return memory[len(nums)-1]
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}
