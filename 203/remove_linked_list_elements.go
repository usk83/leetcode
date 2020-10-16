// 203. Remove Linked List Elements
// https://leetcode.com/problems/remove-linked-list-elements/
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

/*
 * v2. Prepare temporary node and iterate all nodes
 */
func removeElements(head *ListNode, val int) *ListNode {
	head = &ListNode{Next: head}
	for node := head; node != nil; node = node.Next {
		for node.Next != nil && node.Next.Val == val {
			node.Next = node.Next.Next
		}
	}
	return head.Next
}

/*
 * v1. Finding head and iterate all nodes
 */
// func removeElements(head *ListNode, val int) *ListNode {
// 	for head != nil && head.Val == val {
// 		head = head.Next
// 	}
// 	for node := head; node != nil; node = node.Next {
// 		for node.Next != nil && node.Next.Val == val {
// 			node.Next = node.Next.Next
// 		}
// 	}
// 	return head
// }

/*
 * The first solution
 */
// func removeElements(head *ListNode, val int) *ListNode {
// 	for head != nil && head.Val == val {
// 		head = head.Next
// 	}

// 	if head == nil {
// 		return nil
// 	}

// 	node := head
// 	for {
// 		for node.Next != nil && node.Next.Val == val {
// 			node.Next = node.Next.Next
// 		}
// 		node = node.Next
// 		if node == nil {
// 			break
// 		}
// 	}

// 	return head
// }
