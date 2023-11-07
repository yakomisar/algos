package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	tmp := head
	memory := make(map[*ListNode]bool)
	for tmp != nil {
		if _, ok := memory[tmp]; ok {
			return true
		} else {
			memory[tmp] = true
		}
		tmp = tmp.Next
	}
	return false
}

func main() {
	head := ListNode{3, nil}
	one := ListNode{2, nil}
	two := ListNode{0, nil}
	three := ListNode{-4, nil}
	head.Next = &one
	one.Next = &two
	three.Next = &one
	// fmt.Printf("%#v", head)
	fmt.Println(hasCycle(&head))
}
