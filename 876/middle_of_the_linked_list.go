// 876. Middle of the Linked List
// https://leetcode.com/problems/middle-of-the-linked-list/
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
 * v2. Convert the linked list to an array
 * - Time Complexity: O(N)
 * - Space Complexity: O(N) (O(100))
 */
// func middleNode(node *ListNode) *ListNode {
// 	list := make([]*ListNode, 0, 100)
// 	for node != nil {
// 		list = append(list, node)
// 		node = node.Next
// 	}
// 	return list[len(list)>>1]
// }

/*
 * v1. Two pointers
 * - Time Complexity: O(N)
 * - Space Complexity: O(1)
 */
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}
