// 687. Longest Univalue Path
// https://leetcode.com/problems/longest-univalue-path/
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(longestUnivaluePath(BuildTreeNode("[5,4,5,1,1,5]")))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(longestUnivaluePath(BuildTreeNode("[1,4,5,4,4,5]")))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(longestUnivaluePath(BuildTreeNode("[1,null,1,1,1,1,1,1]")))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(longestUnivaluePath(BuildTreeNode("[1,2,3,4]")))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(longestUnivaluePath(BuildTreeNode("[5,4,5,4,4,5,3,4,4,null,null,null,4,null,null,4,null,null,4,null,4,4,null,null,4,4]")))
	pp.Println(6)
	pp.Println("=========================================")
	pp.Println(longestUnivaluePath(BuildTreeNode("[1,2,3,4,2]")))
	pp.Println(1)
	pp.Println("=========================================")
}

// Note: The given binary tree has not more than 10000 nodes.
// The height of the tree is not more than 1000.
func longestUnivaluePath(root *TreeNode) int {
	var longestLength int

	var trace func(node *TreeNode) int
	trace = func(node *TreeNode) int {
		if node == nil || (node.Left == nil && node.Right == nil) {
			return 0
		}

		leftLength := trace(node.Left)
		rightLength := trace(node.Right)

		if node.Left != nil && node.Val == node.Left.Val {
			leftLength++
		} else {
			if leftLength > longestLength {
				longestLength = leftLength
			}
			leftLength = 0
		}

		if node.Right != nil && node.Val == node.Right.Val {
			rightLength++
		} else {
			if rightLength > longestLength {
				longestLength = rightLength
			}
			rightLength = 0
		}

		if leftLength != 0 && rightLength != 0 {
			if currentLongest := leftLength + rightLength; currentLongest > longestLength {
				longestLength = currentLongest
			}
		}

		if leftLength > rightLength {
			return leftLength
		}
		return rightLength
	}

	if length := trace(root); length > longestLength {
		longestLength = length
	}

	return longestLength
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
