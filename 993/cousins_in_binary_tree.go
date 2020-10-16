// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day X
// https://leetcode.com/XXX
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

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	depths := make([][]int, 0, 2)
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, parent, depth int) {
		if node == nil || len(depths) == 2 {
			return
		}
		if node.Val == x || node.Val == y {
			depths = append(depths, []int{parent, depth})
		}
		dfs(node.Left, node.Val, depth+1)
		dfs(node.Right, node.Val, depth+1)
	}
	dfs(root, root.Val, 0)
	return len(depths) == 2 && depths[0][0] != depths[1][0] && depths[0][1] == depths[1][1]
}
