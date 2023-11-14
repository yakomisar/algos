package main

import "fmt"

func isSubsequence(s string, t string) bool {
	left := 0
	right := 0
	for left < len(s) && right < len(t) {
		if s[left] == t[right] {
			left++
			right++
		} else {
			right++
		}
	}
	if len(s) == left {
		return true
	}
	return false
}

func main() {
	s := "axc"
	t := "ahbgc"
	fmt.Println(isSubsequence(s, t))
}
