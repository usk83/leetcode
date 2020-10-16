// 5526. Maximum Number of Achievable Transfer Requests
// https://leetcode.com/contest/weekly-contest-208/problems/maximum-number-of-achievable-transfer-requests/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
}

// 1 <= n <= 20
// 1 <= requests.length <= 16
// requests[i].length == 2
// 0 <= fromi, toi < n

var edges [][]int // 行き先
var graph [][]int // 地点ごとにedge番号

func maximumRequests(n int, requests [][]int) int {
	// 方向付きグラフを作成して
	// 循環したら成立
	// より大きい循環を成立させる必要がある
	// 	backtrakingでつべ手の経路を検証
	// 	何らかの形で成立した場合、その場合の計算をする
	// 成立数が最大になったものを返却する

	edges = requests
	graph = make([][]int, n)
	for i := range graph {
		graph[i] = []int{}
	}

	for i, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], i)
	}

	return f(0, edges[0][0], edges[0][0], 0, 0)
}

func f(startEdge, startNum, currentNum, used, count int) int {
	var maxCount int
	if count != 0 && startNum == currentNum {
		// ここまで成立
	} else {
		// 現在位置から処理

	}
	return maxCount
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
