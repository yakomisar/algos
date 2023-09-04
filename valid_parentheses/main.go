package main

import (
	"fmt"
)

func isValid(s string) bool {
	order := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	result := []string{}
	for i := range s {
		char := string(s[i])
		if val, ok := order[char]; ok {
			if len(result) > 0 {
				if val == result[len(result)-1] {
					result = result[:len(result)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		} else {
			result = append(result, char)
		}

	}
	return len(result) == 0
}

func main() {
	s0 := "]"
	fmt.Println(isValid(s0))

	s1 := "()[]{}"
	fmt.Println(isValid(s1))

	s2 := "(]"
	fmt.Println(isValid(s2))

	s3 := "()"
	fmt.Println(isValid(s3))

	s4 := "{[]}"
	fmt.Println(isValid(s4))

	s5 := "({[)"
	fmt.Println(isValid(s5))
}
