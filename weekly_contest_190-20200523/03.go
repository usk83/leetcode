// 5418. Pseudo-Palindromic Paths in a Binary Tree
// https://leetcode.com/contest/weekly-contest-190/problems/pseudo-palindromic-paths-in-a-binary-tree/
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

func pseudoPalindromicPaths(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	var count int
	freq := make([]int, 10)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		// fmt.Println("=== DEBUG START ======================================")
		// fmt.Printf("current: %d, parent: %d, grandparent: %d\n", node.Val, p, gp)
		// fmt.Println("=== DEBUG END ========================================")

		// if node.Val == p || node.Val == gp {
		// 	yes = true
		// }

		freq[node.Val]++

		if node.Left == nil && node.Right == nil {
			// fmt.Println("=== DEBUG START ======================================")
			// fmt.Println(freq)
			// fmt.Println("=== DEBUG END ========================================")
			var odd int
			for _, f := range freq {
				if f&1 == 1 {
					odd++
				}
				if odd > 1 {
					freq[node.Val]--
					return
				}
			}
			count++
			// if yes {
			// 	count++
			// }
			freq[node.Val]--
			return
		}

		if node.Left != nil {
			dfs(node.Left)
		}

		if node.Right != nil {
			dfs(node.Right)
		}

		freq[node.Val]--

	}
	dfs(root)
	return count
}
