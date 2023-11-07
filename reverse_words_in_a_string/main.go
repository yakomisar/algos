package main

import "fmt"

func isSpace(check byte) bool {
	if check == 32 {
		return true
	}
	return false
}

func trimLeft(s string) int {
	trim := 0
	for _, val := range s {
		if !isSpace(byte(val)) {
			break
		}
		trim++
	}
	return trim
}

func trimRight(s string) int {
	trim := len(s) - 1
	for trim >= 0 {
		if !isSpace(byte(s[trim])) {
			break
		}
		trim--
	}
	return trim
}

func reverseWords(s string) string {
	l := trimLeft(s)
	r := trimRight(s)
	finish := r
	res := ""
	for l <= r {
		if !isSpace(s[r]) {
			r--
			continue
		} else if isSpace(s[r]) {
			if !isSpace(s[r+1]) {
				res += s[r+1 : finish+1]
				res += " "
				finish = r - 1
			} else {
				finish--
			}
		}
		r--
	}
	res += s[r+1 : finish+1]
	return res
}

func main() {
	s := "a good   example"
	fmt.Println(reverseWords(s))
}
