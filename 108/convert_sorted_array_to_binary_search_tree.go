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
func sortedArrayToBST(nums []int) *TreeNode {
	// 真ん中をrootに
	// これを繰り返す

	var convert func(*TreeNode, int, int) *TreeNode
	convert = func(node *TreeNode, start, end int) *TreeNode {
		if start == end {
			return node
		}
		middle := (end + start) / 2

	}

	return convert(nil, 0, len(nums))
}
