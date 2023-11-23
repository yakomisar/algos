package main

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	result := []string{}
	start := 0
	for i := 1; i < len(nums); i++ {
		// Check if the current element is not consecutive
		if nums[i] != nums[i-1]+1 {
			if start == i-1 {
				// Single element range
				result = append(result, fmt.Sprintf("%d", nums[start]))
			} else {
				// Multi-element range
				result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[i-1]))
			}
			start = i
		}
	}
	// Handle the last range
	if start == len(nums)-1 {
		result = append(result, fmt.Sprintf("%d", nums[start]))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[len(nums)-1]))
	}
	return result
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
}
