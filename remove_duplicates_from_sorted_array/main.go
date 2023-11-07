package main

import "fmt"

// Alternative decision
// l := 0
//
//	for r := 1; r < len(nums); r++ {
//		if nums[l] != nums[r] {
//			l++
//			nums[l] = nums[r]
//		}
//	}
//
// return l + 1
func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		} else if nums[i] == nums[i-1] {
			copy(nums[i:], nums[i+1:])
			nums[len(nums)-1] = 0
			nums = nums[:len(nums)-1]
			i--
		}
	}
	return len(nums)
}

func main() {
	nums := []int{1, 1}
	fmt.Println(removeDuplicates(nums))
}
