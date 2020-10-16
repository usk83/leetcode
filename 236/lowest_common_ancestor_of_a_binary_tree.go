// 236. Lowest Common Ancestor of a Binary Tree
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
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

	// belows are wrong
	// Left  *ListNode
	// Right *ListNode
}

/*
 * v2. ツリーが対象を含むかどうかを判定する解法
 * > Runtime: 8 ms, faster than 99.49%
 * > Memory Usage: 7.1 MB, less than 100.00%
 */
// func lowestCommonAncestor(node, p, q *TreeNode) *TreeNode {
// 	if node == nil || node == p || node == q {
// 		return node
// 	}
// 	left, right := lowestCommonAncestor(node.Left, p, q), lowestCommonAncestor(node.Right, p, q)
// 	if left != nil {
// 		if right != nil {
// 			return node
// 		}
// 		return left
// 	}
// 	return right
// }

/*
 * v1. 一致した数を子から親に伝達していく解法
 * > Runtime: 8 ms, faster than 99.49%
 * > Memory Usage: 7.3 MB, less than 100.00%
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(*TreeNode) (*TreeNode, int)
	dfs = func(node *TreeNode) (*TreeNode, int) {
		var count int
		if node == p || node == q {
			count++
		}
		for _, n := range []*TreeNode{node.Left, node.Right} {
			if n == nil {
				continue
			}
			if lca, c := dfs(n); lca != nil {
				return lca, 0
			} else if count += c; count == 2 {
				return node, 0
			}
		}
		return nil, count
	}
	lca, _ := dfs(root)
	return lca
}

/*
 * 再帰なしの実装試作
 *   - DFSでNodeの探索は完了
 *   - 発見した情報を伝達の実装未完
 */
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	node := root
// 	stack := []*TreeNode{}
// 	for length := len(stack); node != nil || length != 0; length = len(stack) {
// 		if node == nil {
// 			node, stack = stack[length-1], stack[:length-1]
// 			continue
// 		}
// 		stack = append(stack, node.Right)
// 		node = node.Left
// 	}
// 	return nil
// }

/*
 * 初期実装
 */
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	ancestorsOfP := []*TreeNode{}
// 	ancestorsOfQ := []*TreeNode{}
//
// 	var dfsP func(node *TreeNode) bool
// 	dfsP = func(node *TreeNode) bool {
// 		if node == nil {
// 			return false
// 		}
// 		if node.Val == p.Val {
// 			ancestorsOfP = append(ancestorsOfP, node)
// 			return true
// 		}
// 		if ok := dfsP(node.Left); ok {
// 			ancestorsOfP = append(ancestorsOfP, node)
// 			return true
// 		}
// 		if ok := dfsP(node.Right); ok {
// 			ancestorsOfP = append(ancestorsOfP, node)
// 			return true
// 		}
// 		return false
// 	}
// 	var dfsQ func(node *TreeNode) bool
// 	dfsQ = func(node *TreeNode) bool {
// 		if node == nil {
// 			return false
// 		}
// 		if node.Val == q.Val {
// 			ancestorsOfQ = append(ancestorsOfQ, node)
// 			return true
// 		}
// 		if ok := dfsQ(node.Left); ok {
// 			ancestorsOfQ = append(ancestorsOfQ, node)
// 			return true
// 		}
// 		if ok := dfsQ(node.Right); ok {
// 			ancestorsOfQ = append(ancestorsOfQ, node)
// 			return true
// 		}
// 		return false
// 	}
//
// 	_ = dfsP(root)
// 	_ = dfsQ(root)
//
// 	ancestorsMap := map[int]bool{}
// 	for _, ancestor := range ancestorsOfP {
// 		ancestorsMap[ancestor.Val] = true
// 	}
//
// 	for _, ancestor := range ancestorsOfQ {
// 		if ancestorsMap[ancestor.Val] {
// 			return ancestor
// 		}
// 	}
//
// 	return nil
// }
