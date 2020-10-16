// 141. Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println("No local tests available.")
	pp.Println("=========================================")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(node *ListNode) bool {
	for fast := node; fast != nil && fast.Next != nil; {
		if fast, node = fast.Next.Next, node.Next; fast == node {
			return true
		}
	}
	return false
}
