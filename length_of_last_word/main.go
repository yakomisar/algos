package main

import "fmt"

func lengthOfLastWord(s string) int {
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 32 && res == 0 {
			continue
		}
		if s[i] != 32 {
			res += 1
		} else {
			break
		}
	}
	return res
}

func main() {
	s := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s))
}
