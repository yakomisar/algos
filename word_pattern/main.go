package main

import "fmt"

func wordPattern(pattern string, s string) bool {
	arr := []string{}
	word := ""
	for _, val := range s {
		if val == 32 && len(word) == 0 {
			continue
		} else if val == 32 {
			arr = append(arr, word)
			word = ""
		} else if len(word) > 0 {
			word += string(val)
		}
	}
	if len(pattern) != len(s) {
		return false
	}
	check := make(map[rune]string)
	for key, val := range pattern {
		if w, ok := check[val]; ok {
			if w != arr[key] {
				return false
			}
		}
	}
	return true
}

func main() {
	pattern := "abba"
	s := "dog cat cat dog"
	fmt.Println(wordPattern(pattern, s))
}
