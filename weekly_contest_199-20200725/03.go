// x. xxx
// https://leetcode.com/problems/xxx/
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

func countPairs(root *TreeNode, distance int) int {
	var count int
	var dfs func(*TreeNode) (ret [11]int)
	dfs = func(node *TreeNode) (ret [11]int) {
		if node != nil {
			if node.Left != nil || node.Right != nil {
				leftNodes := dfs(node.Left)
				rightNodes := dfs(node.Right)

				for l := 1; l < distance; l++ {
					for r := 1; r <= distance-l; r++ {
						count += leftNodes[l] * rightNodes[r]
					}
				}

				for i := 10; i > 0; i-- {
					ret[i] = leftNodes[i-1] + rightNodes[i-1]
				}
			} else {
				// leaf
				ret[1] = 1
			}
		}
		return ret
	}
	dfs(root)
	return count
}
