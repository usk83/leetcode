// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{false, false, true, false, true, true, false}))
	pp.Println(8)
	pp.Println("=========================================")
	pp.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{false, false, true, false, false, true, false}))
	pp.Println(6)
	pp.Println("=========================================")
	pp.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{false, false, false, false, false, false, false}))
	pp.Println(0)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	// 各ノードごとにこの配列を用意
	// pp.Println("=== DEBUG START ======================================")
	tree := make([][]int, n)
	for i := range tree {
		tree[i] = []int{}
	}
	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
	}

	// fmt.Println(tree)

	// 各ノードごとに自分含む子がりんごをもっているかを定義する
	mustVisit := make([]bool, n)
	var dfs1 func(int) bool
	dfs1 = func(node int) bool {
		// この事情がを調査
		// 自分の事情を調査
		// どっちかがtrueならtrueを保存
		// 自分の結果を返却
		var childrenHaveApple bool
		for _, child := range tree[node] {
			if dfs1(child) {
				childrenHaveApple = true
			}
		}
		mustVisit[node] = childrenHaveApple || hasApple[node]
		return mustVisit[node]
	}
	dfs1(0)
	// fmt.Println(mustVisit)
	// pp.Println("=== DEBUG END ========================================")
	// DFSする
	var dfs2 func(int, int) int
	dfs2 = func(node, step int) int {
		for _, child := range tree[node] {
			if mustVisit[child] {
				step = dfs2(child, step+1)
			}
		}
		return step + 1
	}
	return dfs2(0, -1)
}
