// 872. Leaf-Similar Trees
// https://leetcode.com/problems/leaf-similar-trees/
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(leafSimilar(
		BuildTreeNode("[3,5,1,6,2,9,8,null,null,7,4]"),
		BuildTreeNode("[3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]"),
	))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(leafSimilar(
		BuildTreeNode("[3,5,1,null,7,null,2,null,null,null,1,null,1,null,1,null,1,null,1,null,1]"),
		BuildTreeNode("[3,5,1,null,7,null,2,null,null,null,1,null,1,null,1,null,1,null,1,null,1]"),
	))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(leafSimilar(
		BuildTreeNode("[14,86,67,54,33,23,null,75,57,78,null,104,17,7,56,86,109,95,null,null,null,118,null,null,56,18,85,20,38,107,17,40,null,41,122,null,28,null,67,45,null,6,57,null,100,4,31,68,4,88,52,39,96,39,94,4,null,94,8,26,null,95,122,106,56,55,79,4,38,14,1,34,112,30,33,53,60,19,76,null,null,40,90,null,null,111,59,null,97,115,18,27,1,114,100,10,56,91,55,null,null,108,null,115,null,45,null,37,106,29,null,6,47,89,118,96,5,118,92,29,92,100,106,null,87,109,77,null,null,82,112,null,97,51,null,115,null,null,2,118,null,112,121,45,44,35,100,null,51,1,123,57,53,null,null,15,108,66,23,null,null,32,116,null,120,95,null,72,null,56,8,17,95,2,33,null,98,98,69,120,null,null,null,null,null,80,119,42,null,111,118,null,null,94,94,4,null,null,121,49,112]"),
		BuildTreeNode("[29,23,85,84,66,null,61,99,112,88,null,90,87,null,6,16,111,46,null,50,84,113,46,34,76,104,105,null,8,26,58,86,9,47,114,78,58,115,2,59,null,7,107,55,89,51,109,null,null,98,49,21,null,90,107,null,114,121,34,null,60,81,66,86,39,null,null,null,null,118,36,null,102,null,9,95,6,39,90,87,116,25,85,null,9,38,51,1,89,56,118,59,22,null,118,92,60,null,108,null,51,65,44,82,10,37,78,null,null,null,null,45,11,null,28,109,26,81,110,62,26,null,38,43,23,42,105,116,120,95,72,null,56,17,95,2,33,98,null,null,98,69,120,96,5,null,80,119,42,null,111,null,null,null,null,94,94,4,121,null,49,null,112,77,19,null,null,112,104,39,97,51,null,112,121,null,45,null,44,35,null,100,51,1,123,null,57,53,null,10,null,null,15,108,66,null,null,55,106,null,32]"),
	))
	pp.Println(true)
	pp.Println("=========================================")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// As both of the given trees have nodes at most 100,
	// maximal number of leaves is 50.
	leaves := make([]*int, 50)
	var walk func(*TreeNode, *int) bool
	walk = func(root *TreeNode, index *int) bool {
		if root.Left == nil && root.Right == nil {
			if leaves[*index] != nil {
				if *leaves[*index] != root.Val {
					return false
				}
			} else {
				leaves[*index] = &root.Val
			}
			*index++
			return true
		}

		for _, leaf := range []*TreeNode{root.Left, root.Right} {
			if leaf == nil {
				continue
			}
			if !walk(leaf, index) {
				return false
			}
		}

		return true
	}

	index1, index2 := 0, 0
	walk(root1, &index1)
	return walk(root2, &index2)
}

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
