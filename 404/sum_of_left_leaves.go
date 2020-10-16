// 404. Sum of Left Leaves
// https://leetcode.com/problems/sum-of-left-leaves/
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var dfs func(*TreeNode, bool) int
	dfs = func(node *TreeNode, isLeft bool) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil && isLeft {
			return node.Val
		}
		return dfs(node.Left, true) + dfs(node.Right, false)
	}
	return dfs(root, false)
}
