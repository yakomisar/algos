package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	tmp1, tmp2 := list1, list2
	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val <= tmp2.Val {
			current.Next = tmp1
			tmp1 = tmp1.Next
		} else {
			current.Next = tmp2
			tmp2 = tmp2.Next
		}
		current = current.Next
	}
	if tmp1 == nil {
		current.Next = tmp2
	} else if tmp2 == nil {
		current.Next = tmp1
	}
	return dummy.Next
}

func main() {
	list1 := ListNode{1, nil}
	one := ListNode{2, nil}
	two := ListNode{4, nil}
	list1.Next = &one
	one.Next = &two

	list2 := ListNode{1, nil}
	one2 := ListNode{3, nil}
	two2 := ListNode{4, nil}
	list2.Next = &one2
	one2.Next = &two2

	fmt.Println(mergeTwoLists(&list1, &list2))
}
