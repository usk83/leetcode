// 21. Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := ListNode{}
	current := &ListNode{}
	dummy.Next = current
	for l1 != nil || l2 != nil {
		if l1 != nil {
			current.Next = l1
			current = current.Next
			l1 = l1.Next
		}
		if l2 != nil {
			current.Next = l2
			current = current.Next
			l2 = l2.Next
		}
	}
	return dummy.Next.Next
}
