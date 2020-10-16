// 671. Second Minimum Node In a Binary Tree
// https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/
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

func findSecondMinimumValue(root *TreeNode) int {
	second := -1
	node, queue := (*TreeNode)(nil), []*TreeNode{root}
	for len(queue) != 0 {
		node, queue = queue[0], queue[1:] // dequeue
		if node.Left == nil {             // node.Right is also nil
			continue
		}
		for _, child := range []*TreeNode{node.Left, node.Right} {
			if child.Val == root.Val {
				queue = append(queue, child) // enqueue
			} else if second == -1 || child.Val < second {
				second = child.Val
			}
		}
	}
	return second
}

// func findSecondMinimumValue(root *TreeNode) int {
// 	// ルートじゃない数字が出たら確定
// 	// 左右はより小さい方にだけいけばよい
// 	// 左右のminが同じ時は両方調べる

// 	if root == nil {
// 		return -1
// 	}

// 	second := root.Val
// 	q := []*TreeNode{root}
// 	for len(q) != 0 {
// 		node := q[0]
// 		q = q[1:]
// 		if node.Left == nil { // node.Right is also nil
// 			continue
// 		}

// 		for _, child := range []*TreeNode{node.Left, node.Right} {
// 			if child.Val == root.Val {
// 				q = append(q, child)
// 			} else {
// 				if child.Val < second {
// 					second = child.Val
// 				}

// 			}
// 		}

// 		// if node.Left.Val == root.Val {

// 		// }

// 		// if node.Left.Val == root.Val {

// 		// }

// 		// if node.Left.Val == root.Val && node.Right.Val == root.Val {
// 		// 	q = append(q, node.Left, node.Right)
// 		// } else if node.Left.Val == root.Val || node.Right.Val == root.Val {
// 		// 	if node.Left.Val > node.Right.Val {
// 		// 		return node.Left.Val
// 		// 	}
// 		// 	return node.Right.Val
// 		// } else {
// 		// 	if node.Left.Val < node.Right.Val {
// 		// 		return node.Left.Val
// 		// 	}
// 		// 	return node.Right.Val
// 		// }
// 	}

// 	if second == root.Val {
// 		return -1
// 	}

// 	return second
// }

// // func findSecondMinimumValue(root *TreeNode) int {
// // 	if root == nil || root.Left == nil {
// // 		return -1
// // 	}

// // 	first, second := -1, -1
// // 	var dfs func(node *TreeNode)
// // 	dfs = func(node *TreeNode) {
// // 		if node == nil {
// // 			return
// // 		}
// // 		if first != node.Val {
// // 			if first == -1 || node.Val < first {
// // 				first, node.Val = node.Val, first
// // 			}
// // 			if second == -1 || node.Val < second {
// // 				second, node.Val = node.Val, second
// // 			}
// // 		}
// // 		dfs(node.Left)
// // 		dfs(node.Right)
// // 	}

// // 	dfs(root)

// // 	if first == second {
// // 		return -1
// // 	}

// // 	return second
// // }
