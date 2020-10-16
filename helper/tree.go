package main

// package helper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	// pp.Println("======================================")
	// pp.Println(BuildBinaryTree("[]"))
	// pp.Println("======================================")
	// pp.Println(BuildBinaryTree("[3,9,20,null,null,15,7]"))
	// pp.Println("======================================")
	// pp.Println(BuildBinaryTree("[1,null,2,null,3]"))
	// pp.Println("======================================")
	// pp.Println(BuildBinaryTree("[1,null,2,null,3,1,2,null,null,1]"))
	// pp.Println("======================================")

	BuildBinaryTree("[1,null,2,null,3,1,2,null,null,1]").PrintDebugInfo()
}

type treeNodeDebugInfo struct {
	Pos        int
	LeftCount  int
	RightCount int
	Rank       int
	RankPos    int
}

type TreeNode struct {
	Val       int
	Left      *TreeNode
	Right     *TreeNode
	DebugInfo treeNodeDebugInfo
	pp        *pp.PrettyPrinter
}

func (n *TreeNode) PrintDebugInfo() {
	// pp.SetColorScheme()
	// debugInfo := pp.Sprint("Debug")
	// fmt.Println(debugInfo)
	// pp.SetColorScheme()

	n.pp.Println("Debug")

	// var builder strings.Builder
	// builder.WriteString(pp.Sprintf("TreeNode{ Val: %s,  ", n.Val))
	// if n.Left != nil {
	// 	builder.WriteString(pp.Sprintf("Left: %s,  ", n.Left.Val))
	// } else {
	// 	builder.WriteString(pp.Sprintf("Left: %s,  ", nil))
	// }
	// if n.Right != nil {
	// 	builder.WriteString(pp.Sprintf("Right: %s,  ", n.Right.Val))
	// } else {
	// 	builder.WriteString(pp.Sprintf("Right: %s,  ", nil))
	// }
	// builder.WriteString(pp.Sprintf(
	// 	"Pos: %s,  LeftCount: %s,  RightCount: %s,  Rank: %s,  RankPos: %s }\n",
	// 	n.DebugInfo.Pos,
	// 	n.DebugInfo.LeftCount,
	// 	n.DebugInfo.RightCount,
	// 	n.DebugInfo.Rank,
	// 	n.DebugInfo.RankPos,
	// ))
	// fmt.Printf("%v", builder)
}

/*
 *
 *
 *
 * デバッグ情報をもたせる
 *   DebugPrint
 * プリントできるようにする
 *   String
 *
 *
 *
 */

func BuildBinaryTree(treeString string) *TreeNode {
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

	printer := pp.New()
	printer.SetColorScheme(pp.ColorScheme{
		Bool:            pp.Cyan | pp.Bold,
		Integer:         pp.Blue | pp.Bold,
		Float:           pp.Magenta | pp.Bold,
		String:          pp.Red,
		StringQuotation: pp.Red | pp.Bold,
		EscapedChar:     pp.Magenta | pp.Bold,
		FieldName:       pp.Yellow,
		PointerAdress:   pp.Blue | pp.Bold,
		Nil:             pp.Cyan | pp.Bold,
		Time:            pp.Blue | pp.Bold,
		StructName:      pp.Green,
		ObjectLength:    pp.Blue,
	})
	root.pp = printer

	return root
}
