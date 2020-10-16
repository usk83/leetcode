// 328. Odd Even Linked List
// https://leetcode.com/problems/odd-even-linked-list/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 16
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3331/
func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	pp.Println("=========================================")
	fmt.Print("[ ")
	for node := oddEvenList(head); node != nil; node = node.Next {
		fmt.Printf("%d ", node.Val)
	}
	fmt.Println("]")
	pp.Println("=========================================")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	oddTail, evenHead, evenTail := head, head.Next, head.Next
	for node, isOdd := head.Next.Next, true; node != nil; node, isOdd = node.Next, !isOdd {
		if isOdd {
			oddTail, oddTail.Next = node, node
		} else {
			evenTail, evenTail.Next = node, node
		}
	}
	oddTail.Next, evenTail.Next = evenHead, nil
	return head
}
