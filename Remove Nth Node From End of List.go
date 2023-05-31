package main

//https://leetcode.com/problems/remove-nth-node-from-end-of-list/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := *head
	_, l1, l2 := step(&res, n)
	if l2 == nil && l1 != nil {
		res = *l1
	} else {
		if l2 != nil {
			l2.Next = l1
		} else {
			return nil
		}

	}
	return &res
}

func step(l *ListNode, n int) (int, *ListNode, *ListNode) {
	if l.Next != nil {
		num, l1, l2 := step(l.Next, n)
		if num == n {
			return num + 1, l.Next, nil
		}
		if num == n+1 {
			return num + 1, l1, l
		}
		return num + 1, l1, l2
	} else {
		return 2, nil, nil
	}

}
