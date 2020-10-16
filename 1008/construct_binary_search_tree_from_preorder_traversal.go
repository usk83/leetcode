// 1008. Construct Binary Search Tree from Preorder Traversal
// https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/
package main

import (
	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 24
// https://leetcode.com/explore/featured/card/may-leetcoding-challenge/537/week-4-may-22nd-may-28th/3339/
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

const MaxInt = int(^uint(0) >> 1)

func bstFromPreorder(preorder []int) *TreeNode {
	var index int
	var dfs func(ceiling int) *TreeNode
	dfs = func(ceiling int) *TreeNode {
		if index == len(preorder) {
			return nil
		}

		current := preorder[index]
		if current > ceiling {
			return nil
		}

		index++
		return &TreeNode{
			Val:   current,
			Left:  dfs(current),
			Right: dfs(ceiling),
		}
	}
	return dfs(MaxInt)
}
