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
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// As both of the given trees have nodes at most 100,
	// maximal number of leaves is 50
	leaves := make([]int, 0, 50)
	index := 0

	var walk func(*TreeNode, bool) bool
	walk = func(root *TreeNode, check bool) bool {
		if root.Left == nil && root.Right == nil {
			if check {
				ret := leaves[index] == root.Val
				index++
				return ret
			}
			leaves = append(leaves, root.Val)
			return true
		}

		if root.Left != nil {
			if !walk(root.Left, check) {
				return false
			}
		}
		if root.Right != nil {
			if !walk(root.Right, check) {
				return false
			}
		}

		return true
	}

	_ = walk(root1, false)
	return walk(root2, true)
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
