package main

import "fmt"

func wordPattern(pattern string, s string) bool {
	arr := []string{}
	word := ""
	for _, val := range s {
		if string(val) == " " && len(word) == 0 {
			continue
		} else if string(val) == " " {
			arr = append(arr, word)
			word = ""
			continue
		}
		word += string(val)
	}
	if len(word) != 0 {
		arr = append(arr, word)
	}
	if len(pattern) != len(arr) {
		return false
	}
	check := make(map[rune]string)
	for key, val := range pattern {
		if w, ok := check[val]; ok {
			if w != arr[key] {
				return false
			}
		} else {
			check[val] = arr[key]
		}
	}
	check2 := make(map[string]rune)
	for key, val := range pattern {
		if w, ok := check2[arr[key]]; ok {
			if w != val {
				return false
			}
		} else {
			check2[arr[key]] = val
		}
	}
	return true
}

func main() {
	pattern := "abba"
	s := "dog cat cat dog"
	fmt.Println(wordPattern(pattern, s))
}
