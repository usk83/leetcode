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
	if big > small {
		small, big = big, small
	}

	// if root == nil {
	// 	return 0
	// }

	// if small == big {
	// 	for root != nil && root.Val != small {
	// 		if small < root.Val {
	// 			root = root.Left
	// 		} else {
	// 			root = root.Right
	// 		}
	// 	}
	// 	if root == nil {
	// 		return 0
	// 	}
	// 	return root.Val
	// }

	// var sum int
	// if small != root.Val {
	// 	if small < root.Val {
	// 		sum
	// 	} else {
	// 		sum
	// 	}
	// }

	// smallは
	// 左に行ったときは +
	// 	右に行ったときは
	// bigは

	// var sum int

	// if

	// if

	// if

	// 片方が大きい
	// 	現在の値を含む
	// 両方が大きい or 両方が小さい
	// 	現在の値を含まない

	// return 0
}

func walkSmall(node *TreeNode, val int) int {
	if node == nil {
		return 0
	}
	if node.Val == val {
		return node.val
	}

}
