// 5346. Linked List in Binary Tree
// https://leetcode.com/contest/weekly-contest-178/problems/linked-list-in-binary-tree/
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	var result bool
	var dfs func(links []*ListNode, node *TreeNode)
	dfs = func(links []*ListNode, node *TreeNode) {
		if node == nil {
			return
		}

		// size := len(links)
		// backup := map[int]*ListNode{}

		currentLinks := []*ListNode{}

		for i := range links {
			if links[i] != nil && links[i].Val == node.Val {
				if links[i].Next == nil {
					result = true
				}
				currentLinks = append(currentLinks, links[i].Next)

				// links[i] = links[i].Next
				// backup[i] = links[i]
			}
		}
		if head.Val == node.Val {
			if head.Next == nil {
				result = true
			} else {
				// backup[size] = head.Next
				// size++
				// links = append(links, head.Next)
				currentLinks = append(currentLinks, head.Next)
			}
		}

		dfs(currentLinks, node.Left)
		// links = links[:size]
		// for k, v := range backup {
		// 	links[k] = v
		// }
		dfs(currentLinks, node.Right)

		return
	}
	dfs([]*ListNode{}, root)
	return result
}
