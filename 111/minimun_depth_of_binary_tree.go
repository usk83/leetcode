// 111. Minimum Depth of Binary Tree
// https://leetcode.com/problems/minimum-depth-of-binary-tree/
package main

import "fmt"

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func main() {
	fmt.Println(minDepth(helper.BuildTreeNode("[1,null,2,null,3,1,2,null,null,1]")))
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countDepth(0, root)
}

func countDepth(depth int, node *TreeNode) int {
	depth++

	if node.Left == nil && node.Right == nil {
		return depth
	}

	var leftDepth, rightDepth int
	if node.Left != nil {
		leftDepth = countDepth(depth, node.Left)
	}
	if node.Right != nil {
		rightDepth = countDepth(depth, node.Right)
	}

	if leftDepth == 0 {
		return rightDepth
	}
	if rightDepth == 0 {
		return leftDepth
	}
	if leftDepth < rightDepth {
		return leftDepth
	}
	return rightDepth
}
