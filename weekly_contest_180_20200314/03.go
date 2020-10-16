// 5179. Balance a Binary Search Tree
// https://leetcode.com/contest/weekly-contest-180/problems/balance-a-binary-search-tree/

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	balanceBST(BuildTreeNode("[1,null,2,null,3,null,4,null,null]"))
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
func balanceBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	nums := []int{}
	var dfs1 func(node *TreeNode)
	dfs1 = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs1(node.Left)
		dfs1(node.Right)
		nums = append(nums, node.Val)
	}
	dfs1(root)

	sort.Ints(nums)

	var dfs func(node *TreeNode, start, end int)
	dfs = func(node *TreeNode, start, end int) {
		if start == end {
			return
		}
		mid := (start + end) >> 1
		node.Val = nums[mid]
		if start != mid {
			node.Left = &TreeNode{}
			dfs(node.Left, start, mid)
		}
		if end-1 != mid {
			node.Right = &TreeNode{}
			dfs(node.Right, mid+1, end)
		}
	}
	r := &TreeNode{}
	dfs(r, 0, len(nums))
	return r
}

// [2,1,3,null,null,null,4]

// Input: [1,null,2,null,3,null,4,null,null]
// Output: [3,2,4,1,0,null,0,null,0]
// Expected: [2,1,3,null,null,null,4]

func BuildTreeNode(treeString string) *TreeNode {
	treeNodeStrings := strings.Split(treeString[1:len(treeString)-1], ",")

	var nodeLength = len(treeNodeStrings)

	if nodeLength == 1 && treeNodeStrings[0] == "" || treeNodeStrings[0] == "null" {
		return nil
	}

	rootNodeString := treeNodeStrings[0]
	val, err := strconv.Atoi(rootNodeString)
	if err != nil {
		panic(fmt.Sprintf("数値変換エラー: %s", err.Error()))
	}
	root := &TreeNode{
		Val: val,
	}
	treeNodeStrings = treeNodeStrings[1:len(treeNodeStrings)]

	tips := make([]*TreeNode, (nodeLength+1)/2)
	tips[0] = root
	tipsLength := 1

	for len(treeNodeStrings) != 0 {
		breadth := tipsLength * 2
		var children []string
		if len(treeNodeStrings) < breadth {
			children = treeNodeStrings
			treeNodeStrings = []string{}
		} else {
			children = treeNodeStrings[0:breadth]
			treeNodeStrings = treeNodeStrings[breadth:len(treeNodeStrings)]
		}

		childrenNodes := make([]*TreeNode, breadth)
		for i, v := range children {
			if v == "null" {
				continue
			}
			val, err := strconv.Atoi(v)
			if err != nil {
				panic(fmt.Sprintf("数値変換エラー: %s", err.Error()))
			}
			childrenNodes[i] = &TreeNode{
				Val: val,
			}
		}

		for i := 0; i < breadth/2; i++ {
			left := childrenNodes[2*i]
			right := childrenNodes[2*i+1]
			if tips[i] == nil {
				if left == nil && right == nil {
					continue
				}
				panic("Nil NodeはChildrenを持つことができません")
			}

			tips[i].Left = left
			tips[i].Right = right
		}

		length := 0
		for i := range childrenNodes {
			if childrenNodes[i] == nil {
				continue
			}
			tips[length] = childrenNodes[i]
			length++
		}
		tipsLength = length
	}

	return root
}
