// x. xxx
// https://leetcode.com/problems/xxx/
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
func longestUnivaluePath(root *TreeNode) int {
	var longest int
	var walk func(node *TreeNode) int
	walk = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		var ok bool
		left := walk(node.Left)
		right := walk(node.Right)

		// fmt.Printf("left: %d, right: %d\n", left, right)
		longest = max(longest, max(left, right))

		if node.Left != nil && node.Left.Val == node.Val {
			ok = true
			left++
		} else {
			left = 0
		}
		if node.Right != nil && node.Right.Val == node.Val {
			ok = true
			right++
		} else {
			right = 0
		}

		if ok {
			longest = max(longest, left+right)
			return max(left, right)
		}
		return 0
	}
	// rootLenfth := walk(root)
	_ = walk(root)
	// return max(longest, rootLenfth)
	return longest
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
