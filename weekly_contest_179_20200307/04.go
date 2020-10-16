// 5355. Frog Position After T Seconds
// https://leetcode.com/contest/weekly-contest-179/problems/frog-position-after-t-seconds/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 2, 4))
	pp.Println(0.16666666666666666)
	pp.Println("=========================================")
	pp.Println(frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 1, 7))
	pp.Println(0.3333333333333333)
	pp.Println("=========================================")
	pp.Println(frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 20, 6))
	pp.Println(0.16666666666666666)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

// vertex: 1-n
// start: vertex 1
// 再訪問不可
// ランダムな確率でどこかへ
// 移動先がないときは止まる
// t秒後にtargetに入る確率
func frogPosition(n int, edges [][]int, t int, target int) float64 {
	// Graphを構築する

	graph := make([][]int, n)
	for i := range graph {
		graph[i] = []int{}
	}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	current := map[int][2]int{
		1, 1,
	}

	// 現在の頂点頂点一覧
	// そこまでの確率をもつ
	for sec := 0; sec < t; sec++ {

	}

	// 最終的にtargetについていくつか
	if pos, ok := current[target]; ok {
		return pos[1]
	}

	return 0
}
