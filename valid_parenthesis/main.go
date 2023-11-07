package main

import (
	"fmt"
)

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	stack := []string{}
	rule := map[string]string{")": "(", "}": "{", "]": "["}
	for _, val := range s {
		if len(stack) == 0 {
			stack = append(stack, string(val))
			continue
		}
		if par, ok := rule[string(val)]; ok {
			if stack[len(stack)-1] == par {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, string(val))
		}
	}
	return len(stack) == 0
}

func main() {
	s1 := "()"
	s2 := "()[]{}"
	s3 := "(]"
	fmt.Println(isValid(s1))
	fmt.Println(isValid(s2))
	fmt.Println(isValid(s3))
}
