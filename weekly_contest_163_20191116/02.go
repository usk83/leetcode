package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	findElements := Constructor(BuildTreeNode("[-1,null,-1,-1,null,-1]"))
	pp.Println(findElements.Find(2))
	pp.Println(findElements.Find(3))
	pp.Println(findElements.Find(4))
	pp.Println(findElements.Find(5))
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// rootは0
// If treeNode.val == x and treeNode.left != null, then treeNode.left.val == 2 * x + 1
// If treeNode.val == x and treeNode.right != null, then treeNode.right.val == 2 * x + 2
type FindElements struct {
	root *TreeNode
	dict map[int]struct{}
}

func Constructor(root *TreeNode) FindElements {
	findElements := FindElements{
		root: root,
		dict: map[int]struct{}{},
	}

	findElements.dict[0] = struct{}{}

	root.Val = 0
	findElements.dict[0] = struct{}{}

	var next func(node *TreeNode, val int)
	next = func(node *TreeNode, val int) {
		if node == nil {
			return
		}

		node.Val = val

		findElements.dict[val] = struct{}{}
		next(node.Left, val*2+1)
		next(node.Right, val*2+2)
		return
	}

	next(root.Left, 1)
	next(root.Right, 2)

	return findElements
}

func (this *FindElements) Find(target int) bool {
	_, found := this.dict[target]
	return found
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
