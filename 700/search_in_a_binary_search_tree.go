// 700. Search in a Binary Search Tree
// https://leetcode.com/problems/search-in-a-binary-search-tree/
package main

func main() {}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

// �� �����nil�����å����ʤ��Ƥ�ͨ��

// func searchBST(root *TreeNode, val int) *TreeNode {
// 	if root == nil || root.Val == val {
// 		return root
// 	}
// 	if val < root.Val {
// 		return searchBST(root.Left, val)
// 	}
// 	return searchBST(root.Right, val)
// }
