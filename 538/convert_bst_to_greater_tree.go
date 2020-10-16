// 538. Convert BST to Greater Tree
// https://leetcode.com/problems/convert-bst-to-greater-tree/
package main

func main() {
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

// Version 3
func convertBST(root *TreeNode) *TreeNode {
	_ = modifyTree(root, 0)
	return root
}

func modifyTree(node *TreeNode, sum int) int {
	if node == nil {
		return sum
	}
	node.Val += modifyTree(node.Right, sum)
	return modifyTree(node.Left, node.Val)
}

// Version 2_
// func convertBST(root *TreeNode) *TreeNode {
// 	_ = modifyTree(root, 0)
// 	return root
// }
//
// func modifyTree(node *TreeNode, sum int) int {
// 	if node == nil {
// 		return sum
// 	}
// 	sum = modifyTree(node.Right, sum)
// 	sum += node.Val
// 	node.Val = sum
// 	sum = modifyTree(node.Left, sum)
// 	return sum
// }

// Version 1
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	_ = modifyTree(root, 0)
	return root
}

func modifyTree(node *TreeNode, sum int) int {
	if node.Right != nil {
		sum = modifyTree(node.Right, sum)
	}
	sum += node.Val
	node.Val = sum
	if node.Left != nil {
		sum = modifyTree(node.Left, sum)
	}
	return sum
}
