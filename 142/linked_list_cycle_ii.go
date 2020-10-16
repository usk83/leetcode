// 142. Linked List Cycle II
// https://leetcode.com/problems/linked-list-cycle-ii/
package main

import "github.com/k0kubun/pp"

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
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	visited := map[*ListNode]struct{}{}
	for head != nil {
		if _, found := visited[head]; found {
			return head
		}
		visited[head] = struct{}{}
		head = head.Next
	}
	return nil
}
