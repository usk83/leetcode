// 437. Path Sum III
// https://leetcode.com/problems/path-sum-iii/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println("No local tests available.")
	pp.Println("=========================================")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	var count int
	if node.Val == sum {
		count++
	}
	return count + helper(node.Left, sum-node.Val) + helper(node.Right, sum-node.Val)
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return helper(root, sum) + helper(root.Left, sum) + helper(root.Right, sum)
}

// func pathSum(root *TreeNode, sum int) int {
// 	sums := []int{}
// 	var walk func(node *TreeNode) int
// 	walk = func(node *TreeNode) int {
// 		if node == nil {
// 			return 0
// 		}
//
// 		sums = append(sums, sum)
// 		for i := range sums {
// 			sums[i] -= node.Val
// 		}
//
// 		count := walk(node.Left) + walk(node.Right)
//
// 		for i := range sums {
// 			if sums[i] == 0 {
// 				count++
// 			}
// 			sums[i] += node.Val
// 		}
// 		sums = sums[:len(sums)-1]
//
// 		return count
// 	}
// 	return walk(root)
// }

// func pathSum(root *TreeNode, sum int) int {
// 	count, length, sums := 0, 0, make([]int, 10)
//
// 	var walk func(node *TreeNode)
// 	walk = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
//
// 		if length == len(sums) {
// 			sums = append(sums, 0)
// 		}
// 		length++
//
// 		for i := 0; i < length; i++ {
// 			sums[i] += node.Val
// 			if sums[i] == sum {
// 				count++
// 			}
// 		}
//
// 		walk(node.Left)
// 		walk(node.Right)
//
// 		for i := 0; i < length; i++ {
// 			sums[i] -= node.Val
// 		}
// 		length--
// 	}
//
// 	walk(root)
//
// 	return count
// }
