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

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	var walk func(node *TreeNode) *TreeNode
	walk = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		node.Left = walk(node.Left)
		node.Right = walk(node.Right)

		if node.Left == nil && node.Right == nil && node.Val == target {
			return nil
		}
		return node
	}

	return walk(root)
}
