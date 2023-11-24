package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	res := ""
	for i := range strs[0] {
		for s := range strs {
			if i == len(strs[s]) || strs[s][i] != strs[0][i] {
				return res
			}
		}
		res += string(strs[0][i])
	}
	return res
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
