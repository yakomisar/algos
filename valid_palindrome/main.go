package main

import "fmt"

func isPalindrome(s string) bool {
	res := []byte{}
	check := []byte{}
	for i := range s {
		if s[i] >= 65 && s[i] <= 90 {
			res = append(res, s[i])
		} else if s[i] >= 97 && s[i] <= 122 {
			res = append(res, s[i]-32)
		} else if s[i] >= 48 && s[i] <= 57 {
			res = append(res, s[i])
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= 65 && s[i] <= 90 {
			check = append(check, s[i])
		} else if s[i] >= 97 && s[i] <= 122 {
			check = append(check, s[i]-32)
		} else if s[i] >= 48 && s[i] <= 57 {
			check = append(check, s[i])
		}
	}
	return string(check) == string(res)
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
