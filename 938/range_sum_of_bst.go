// 938. Range Sum of BST
// https://leetcode.com/problems/range-sum-of-bst/
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
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// The number of nodes in the tree is at most 10000.
// The final answer is guaranteed to be less than 2^31.
func rangeSumBST(root *TreeNode, small int, big int) int {
	// if big < small {
	// 	small, big = big, small
	// }

	var walk func(*TreeNode, int val) int
	walk = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		if val == node.Val {
			return val
		}

		var left, right *TreeNode
		if small < node.Val {
			left = node.Left
		} else {
			left = node.Right
		}
		if big < node.Val {
			right = node.Left
		} else {
			right = node.Right
		}

		if small < node.Val && node.Val < big {
			return walk(left) + node.Val + walk(right)
		}

		return walk(left) + walk(right)
	}

	return walk(root, small) + walk(root, big)
}
