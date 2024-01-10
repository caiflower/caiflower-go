package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := new(ListNode)
	l.Val = 9
	l.Next = new(ListNode)
	l.Next.Val = 1
	numbers := addTwoNumbers(l, nil)
	fmt.Printf("%v", numbers)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	carry := 0
	var p *ListNode
	l1 = reversed(l1)
	l2 = reversed(l2)
	for l1 != nil || l2 != nil || carry != 0 {
		res = new(ListNode)
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		res.Val = sum % 10
		carry = sum / 10
		res.Next = p
		p = res
	}
	return res
}

func reversed(l *ListNode) *ListNode {
	if l == nil {
		return l
	}
	var p *ListNode
	for l != nil {
		next := l.Next
		l.Next = p
		p = l
		l = next
	}
	return p
}
