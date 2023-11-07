package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tst := 0
	tmp := x
	for tmp > 0 {
		last_digit := tmp % 10
		tst = (tst * 10) + last_digit
		tmp /= 10
	}
	return tst == x
}

func main() {
	x := 121
	fmt.Println(isPalindrome(x))
	x = 10
	fmt.Println(isPalindrome(x))
}
