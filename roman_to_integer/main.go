package main

import "fmt"

func romanToInt(s string) int {
	dict := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		if i == (len(s) - 1) {
			result += dict[rune(s[i])]
			continue
		} else if s[i] == 'I' && (s[i+1] == 'V' || s[i+1] == 'X') {
			result -= dict[rune(s[i])]
		} else if s[i] == 'X' && (s[i+1] == 'L' || s[i+1] == 'C') {
			result -= dict[rune(s[i])]
		} else if s[i] == 'C' && (s[i+1] == 'D' || s[i+1] == 'M') {
			result -= dict[rune(s[i])]
		} else {
			result += dict[rune(s[i])]
		}
	}
	return result
}

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}
