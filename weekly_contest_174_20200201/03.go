// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxProduct(BuildTreeNode("[1,2,3,4,5,6]")))
	pp.Println(110)
	pp.Println("=========================================")
	// pp.Println(maxProduct(BuildTreeNode("[1,null,2,3,4,null,null,5,6]")))
	// pp.Println(90)
	// pp.Println("=========================================")
	// pp.Println(maxProduct(BuildTreeNode("[2,3,9,10,7,8,6,5,4,11,1]")))
	// pp.Println(1025)
	// pp.Println("=========================================")
	// pp.Println(maxProduct(BuildTreeNode("[1,1]")))
	// pp.Println(1)
	// pp.Println("=========================================")
}

// 2つの値が極力近くなったほうがよい
// dfsでdp
// 切った場合切らなかった場合の両方を保持する

func maxProduct(root *TreeNode) int {
	left, right := 0, 0
	var first func(node *TreeNode)
	first = func(node *TreeNode) {
		if node == nil {
			return
		}

		first(node.Left)
		first(node.Right)

		if left == 0 {
			left = node.Val
		} else {
			right += node.Val
		}
	}
	first(root)

	diff := abs(left, right)

	pp.Println(left, right, diff)

	cl, cr := left, right
	edge := false

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		dfs(node.Right)

		if !edge {
			edge = true
			return
		}

		cl += node.Val
		cr -= node.Val

		pp.Println(cl, cr, abs(cl, cr))

		if d := abs(cl, cr); d < diff {
			diff = d
			left = cl
			right = cr
			pp.Println("update")
		}
	}
	dfs(root)

	return (left * right) % (10 ^ 9 + 7)
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
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
