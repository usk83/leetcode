// 92. Reverse Linked List II
// https://leetcode.com/problems/reverse-linked-list-ii/
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
	Val   int
	Next  *ListNode
	debug struct {
		index int
	}
}

// [1,2,3,4,5], 2, 4
// [5], 1, 1
// [3,5], 1, 2

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	m, n = m-1, n-1
	end := head
	for i := 0; i < m; i++ {
		end = end.Next
	}
	start := end.Next
	if start == nil {
		return head
	}
	current, next := start, start.Next
	for i := m; i < n; i++ {
		if next != nil {
			next.Next, current, next = current, next, next.Next
		} else {
			current = next
		}
	}
	start.Next, end.Next = next, current
	return head
}
