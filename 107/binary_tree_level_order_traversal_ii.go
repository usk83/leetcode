// 107. Binary Tree Level Order Traversal II
// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/
package main

import (
	"github.com/k0kubun/pp"
)

// July LeetCoding Challenge Day 2
// https://leetcode.com/explore/featured/card/july-leetcoding-challenge/544/week-1-july-1st-july-7th/3378/
func main() {
	pp.Println("=========================================")
	pp.Println(levelOrderBottom(nil))
	pp.Println([][]int{})
	pp.Println("=========================================")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	answer := [][]int{}
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(answer) > level {
			answer[level] = append(answer[level], node.Val)
		} else {
			answer = append(answer, []int{node.Val})
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	for left, right := 0, len(answer)-1; left < right; left, right = left+1, right-1 {
		answer[left], answer[right] = answer[right], answer[left]
	}
	return answer
}
