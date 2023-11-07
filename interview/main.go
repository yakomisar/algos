package main

import "fmt"

func main() {
	fruit := make([]string, 5)
	fmt.Println("Fruit:", fruit)

	fruit[0] = "banana"
	fruit[1] = "cucumber"
	fruit[2] = "grape"
	fruit[3] = "watermelon"
	fruit[4] = "pineapple"

	fmt.Println("Fruit:", fruit)
	xs := fruit[1:3:6]

	fmt.Println("XS:", xs)
	fmt.Println("XS:", "LEN:", len(xs), "CAP:", cap(xs))
	xs[0] = "apple"
	xs = append(xs, "mango")

	fmt.Println("XS:", xs)
	fmt.Println("Fruit:", fruit)
}

// func main() {
//     people := make([]person, 2)
//     p1 := &people[1]
//
//     p1.age++
//
//     people = append(people, person{})
//
//     p1.age
//
//     fmt.Println(people[1].age)
//     fmt.Println(p1.age)
// }
