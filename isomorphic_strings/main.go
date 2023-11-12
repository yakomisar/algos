package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	check := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if v, ok := check[t[i]]; ok {
			if v != s[i] {
				return false
			}
		} else {
			check[t[i]] = s[i]
		}
	}
	check2 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if v, ok := check2[s[i]]; ok {
			if v != t[i] {
				return false
			}
		} else {
			check2[s[i]] = t[i]
		}
	}
	return true
}

func main() {
	s := "badc"
	t := "baba"
	fmt.Println(isIsomorphic(s, t))
}
