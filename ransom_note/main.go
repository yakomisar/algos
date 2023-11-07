package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	a := make(map[byte]int)
	for i := range []rune(magazine) {
		a[magazine[i]]++
	}
	for j := range []rune(ransomNote) {
		if _, ok := a[ransomNote[j]]; ok {
			if a[ransomNote[j]] == 0 {
				return false
			} else {
				a[ransomNote[j]]--
			}
		} else {
			return false
		}
	}
	return true
}

func main() {
	ransomNote := "aa"
	magazine := "ab"
	fmt.Println(canConstruct(ransomNote, magazine))
}
