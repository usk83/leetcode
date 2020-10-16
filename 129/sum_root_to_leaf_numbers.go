// 129. Sum Root to Leaf Numbers
// https://leetcode.com/problems/sum-root-to-leaf-numbers/
package main

import (
	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 26
// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/542/week-4-june-22nd-june-28th/3372/
func main() {
	pp.Println("=========================================")
	pp.Println(sumNumbers(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}))
	pp.Println(25)
	pp.Println("=========================================")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 2. Simpler version
 * ref: https://leetcode.com/problems/sum-root-to-leaf-numbers/discuss/41363/Short-Java-solution.-Recursion.
 */
// func sumNumbers(root *TreeNode) int {
// 	var dfs func(node *TreeNode, num int) int
// 	dfs = func(node *TreeNode, num int) int {
// 		if node == nil {
// 			return 0
// 		}
// 		num = num*10 + node.Val
// 		if node.Left == nil && node.Right == nil {
// 			return num
// 		}
// 		return dfs(node.Left, num) + dfs(node.Right, num)
// 	}
// 	return dfs(root, 0)
// }

/**
 * 1. The first solution
 */
func sumNumbers(root *TreeNode) int {
	var sum int
	var dfs func(node *TreeNode, num int) bool
	dfs = func(node *TreeNode, num int) bool {
		if node == nil {
			return true
		}
		num += node.Val
		isLeaf := dfs(node.Left, num*10)
		isLeaf = dfs(node.Right, num*10) && isLeaf
		if isLeaf {
			sum += num
		}
		return false
	}
	dfs(root, 0)
	return sum
}
