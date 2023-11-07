package main

import "fmt"

func intToRoman(num int) string {
	arr := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	hashMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	res := ""
	for i := len(arr) - 1; i >= 0; i-- {
		count := num / arr[i]
		if count != 0 {
			for count > 0 {
				res += hashMap[arr[i]]
				count--
			}
			num %= arr[i]
		}
	}
	return res
}

func main() {
	num := 3
	fmt.Println(intToRoman(num))
}
