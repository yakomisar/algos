package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] {
			substring := haystack[i : len(needle)+i]
			if substring == needle {
				return i
			}
		}
	}
	return -1
}

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
}
