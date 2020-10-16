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
}

// vertex: 1-n
// start: vertex 1
// 再訪問不可
// ランダムな確率でどこかへ
// 移動先がないときは止まる
// t秒後にtargetに入る確率
func frogPosition(n int, edges [][]int, t int, target int) float64 {
	// Graphを構築する
	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = []int{}
	}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	type status struct {
		pos     float64
		visited map[int]bool
	}

	current := map[int][]status{
		1: []status{
			status{
				pos: 1.0,
				visited: map[int]bool{
					1: true,
				},
			},
		},
	}

	// 現在の頂点頂点一覧
	// そこまでの確率をもつ
	for sec := 0; sec < t; sec++ {
		next := map[int][]status{}
		for vertex, stts := range current {
			for _, stt := range stts {

				stt.visited[vertex] = true

				ways := []int{}
				for _, dist := range graph[vertex] {
					if !stt.visited[dist] {
						ways = append(ways, dist)
					}
				}
				if len(ways) == 0 {
					if _, ok := next[vertex]; ok {
						next[vertex] = append(next[vertex], status{
							pos:     stt.pos,
							visited: stt.visited,
						})
					} else {
						next[vertex] = []status{
							status{
								pos:     stt.pos,
								visited: stt.visited,
							},
						}
					}
				}

				for _, way := range ways {
					if _, ok := next[way]; ok {
						next[way] = append(next[way], status{
							pos:     stt.pos / float64(len(ways)),
							visited: stt.visited,
						})
					} else {
						next[way] = []status{
							status{
								pos:     stt.pos / float64(len(ways)),
								visited: stt.visited,
							},
						}
					}
				}

			}
		}
		current = next
	}

	// 最終的にtargetについていくつか
	if result, ok := current[target]; ok {
		pos := 0.0
		for _, stt := range result {
			pos += stt.pos
		}
		return pos
	}

	return 0.0
}
