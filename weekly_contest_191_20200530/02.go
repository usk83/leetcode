// 5425. Maximum Area of a Piece of Cake After Horizontal and Vertical Cuts
// https://leetcode.com/contest/weekly-contest-191/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxArea(5, 4, []int{1, 2, 4}, []int{1, 3}))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(maxArea(5, 4, []int{3, 1}, []int{1}))
	pp.Println(6)
	pp.Println("=========================================")
	pp.Println(maxArea(5, 4, []int{3}, []int{3}))
	pp.Println(9)
	pp.Println("=========================================")
}

const mod = 1000000007

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	// 最初と最後考慮
	heightDiff, widthDiff := horizontalCuts[0], verticalCuts[0]
	for i := range horizontalCuts[1:] {
		heightDiff = max(heightDiff, horizontalCuts[i+1]-horizontalCuts[i])
	}
	heightDiff = max(heightDiff, h-horizontalCuts[len(horizontalCuts)-1])

	for i := range verticalCuts[1:] {
		widthDiff = max(widthDiff, verticalCuts[i+1]-verticalCuts[i])
	}
	widthDiff = max(widthDiff, w-verticalCuts[len(verticalCuts)-1])

	return heightDiff % mod * widthDiff % mod
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
