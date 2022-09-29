package main

import "fmt"

//https://leetcode.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := *l1
	tmp1 := &res
	tmp2 := &*l2
	next := 0
	for {

		if tmp1.Next == nil && tmp2.Next != nil {
			tmp1.Next = tmp2.Next
			tmp2.Next = nil
		}
		tmp1.Val += tmp2.Val + next
		tmp2.Val = 0
		if tmp1.Val >= 10 {
			tmp1.Val = tmp1.Val % 10
			next = 1
		} else {
			next = 0
		}
		if tmp1.Next == nil && next != 0 {
			t := ListNode{next, nil}
			tmp1.Next = &t
			next = 0
		}
		if tmp2.Next == nil && tmp1.Next == nil {
			break
		}
		if tmp1.Next != nil {
			tmp1 = tmp1.Next
		}
		if tmp2.Next != nil {
			tmp2 = tmp2.Next
		}
	}
	return &res
}

func main() {
	var res, tmp ListNode
	fmt.Println(res, tmp)
}
