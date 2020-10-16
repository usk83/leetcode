// 112. Path Sum
// https://leetcode.com/problems/path-sum/
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

// func hasPathSum(root *TreeNode, sum int) bool {
// 	stack := []*TreeNode{root}
// 	length := 1
// 	push := func(child *TreeNode) {
// 		if len(stack) > length {
// 			stack[length] = child
// 		} else {
// 			stack = append(stack, child)
// 		}
// 		length++
// 	}
// 	pop := func() *TreeNode {
// 		length--
// 		return stack[length]
// 	}

// 	curSum := 0
// 	for length != 0 {
// 		node := pop()
// 		if node == nil {
// 			continue
// 		}
// 		curSum += node.Val
// 		if node.Left == nil && node.Right == nil {
// 			if curSum == sum {
// 				return true
// 			}
// 			curSum -= node.Val
// 			continue
// 		}
// 		push(node.Right)
// 		push(node.Left)
// 	}

// 	return false
// }

func hasPathSum(node *TreeNode, sum int) bool {
	if node == nil {
		return false
	}
	curSum := sum - node.Val
	if node.Left == nil && node.Right == nil {
		return curSum == 0
	}
	return hasPathSum(node.Left, curSum) || hasPathSum(node.Right, curSum)
}

// func hasPathSum(root *TreeNode, sum int) bool {
// 	var dfs func(*TreeNode, int) bool
// 	dfs = func(node *TreeNode, prevSum int) bool {
// 		if node == nil {
// 			return false
// 		}
// 		curSum := prevSum + node.Val
// 		if node.Left == nil && node.Right == nil {
// 			return curSum == sum
// 		}
// 		return dfs(node.Left, curSum) || dfs(node.Right, curSum)
// 	}
// 	return dfs(root, 0)
// }
