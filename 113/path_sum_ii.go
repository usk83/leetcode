// 113. Path Sum II
// https://leetcode.com/problems/path-sum-ii/
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

func traverse(node *TreeNode, currentPath []int, rest int, acc [][]int) [][]int {
	if node == nil {
		return acc
	}
	rest -= node.Val
	currentPath = append(currentPath, node.Val)
	if node.Left == nil && node.Right == nil {
		if rest != 0 {
			return acc
		}
		path := make([]int, len(currentPath))
		copy(path, currentPath)
		return append(acc, path)
	}
	return traverse(node.Right, currentPath, rest, traverse(node.Left, currentPath, rest, acc))
}

func pathSum(root *TreeNode, sum int) [][]int {
	return traverse(root, []int{}, sum, [][]int{})
}

/*
 * the first solution
 */
// func pathSum(root *TreeNode, sum int) [][]int {
// 	paths := [][]int{}
// 	var dfs func(node *TreeNode, path []int, reft int)
// 	dfs = func(node *TreeNode, path []int, reft int) {
// 		if node == nil {
// 			return
// 		}
// 		reft -= node.Val
// 		path = append(path, node.Val)
// 		if node.Left == nil && node.Right == nil {
// 			if reft == 0 {
// 				answer := []int{}
// 				for _, p := range path {
// 					answer = append(answer, p)
// 				}
// 				paths = append(paths, answer)
// 			}
// 			return
// 		}
// 		dfs(node.Left, path, reft)
// 		dfs(node.Right, path, reft)
// 	}
// 	dfs(root, []int{}, sum)
// 	return paths
// }
