// 222. Count Complete Tree Nodes
// https://leetcode.com/problems/count-complete-tree-nodes/
package main

import (
	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 23
// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/542/week-4-june-22nd-june-28th/3369/
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

/**
 * 4. 左右見て完全二分木ではない方に進んでいく方法
 */
func countNodes(node *TreeNode) int {
	if node == nil {
		return 0
	}

	// 完全二分木であるかを確認

	// return 1 + countNodes([]int{}node.Left) + countNodes(node.Right)
}

/**
 * 3. 高さを算出して、左右みて比較していく解法
 */
// func countNodes(root *TreeNode) int {
// }

/**
 * 2. Binary Search
 * - Time Complexity: O(logN^2)
 * - Space Complexity: O(1)
 */
// func countNodes(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
//
// 	// 1. Find out the last rank
// 	var rank uint
// 	for pointer := root; pointer.Left != nil; pointer = pointer.Left {
// 		rank++
// 	}
//
// 	// 2. set left and right
// 	left, right := 2, 1<<rank+1
// 	for left < right {
// 		// 3. while left < right, take a mid and check to update left or right
// 		if mid := (left + right) >> 1; isExist(root, mid, rank) {
// 			left = mid + 1
// 		} else {
// 			right = mid
// 		}
// 	}
//
// 	// 4. calcurate the number of nodes
// 	return (1 << rank) - 1 + left - 1
// }
//
// func isExist(node *TreeNode, target int, rank uint) bool {
// 	for rank > 0 {
// 		if mid := 1 << rank >> 1; target <= mid {
// 			node = node.Left
// 		} else {
// 			node = node.Right
// 			target -= mid
// 		}
// 		rank--
// 	}
// 	return node != nil
// }

/**
 * 1. the first solution - DFS
 * - Time Complexity: O(N)
 * - Space Complexity: O(logN)
 */
// func countNodes(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	leavesRank, count := -1, 0
// 	var dfs func(*TreeNode, int) bool
// 	dfs = func(node *TreeNode, rank int) bool {
// 		if leavesRank < 0 {
// 			if node.Left != nil {
// 				count |= 1 << rank
// 				if dfs(node.Left, rank+1) {
// 					return true
// 				}
// 			} else {
// 				leavesRank = rank
// 				count++
// 				return false
// 			}
// 		} else {
// 			if rank == leavesRank-1 {
// 				if node.Left != nil {
// 					count++
// 				} else {
// 					return true
// 				}
// 			} else {
// 				if node.Left != nil {
// 					if dfs(node.Left, rank+1) {
// 						return true
// 					}
// 				} else {
// 					return true
// 				}
// 			}
// 		}
// 		if rank == leavesRank-1 {
// 			if node.Right != nil {
// 				count++
// 				return false
// 			} else {
// 				return true
// 			}
// 		}
// 		if node.Right != nil {
// 			return dfs(node.Right, rank+1)
// 		}
// 		return true
// 	}
// 	dfs(root, 0)
// 	return count
// }
