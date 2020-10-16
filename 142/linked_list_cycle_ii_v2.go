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
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	for head != fast {
		head, fast = head.Next, fast.Next
	}

	return head
}
