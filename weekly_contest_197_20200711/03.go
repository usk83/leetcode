// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxProbability(5, [][]int{{1, 4}, {2, 4}, {0, 4}, {0, 3}, {0, 2}, {2, 3}}, []float64{0.37, 0.17, 0.93, 0.23, 0.39, 0.04}, 3, 4))
	pp.Println("0.2XXXXXXXXXXXX")
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	type edge struct {
		to   int
		prob float64
	}

	graph := make([][]edge, n)
	for i, e := range edges {
		graph[e[0]] = append(graph[e[0]], edge{e[1], succProb[i]})
		graph[e[1]] = append(graph[e[1]], edge{e[0], succProb[i]})
	}

	probs := make([]float64, n)
	probs[start] = 1.0

	queue := []int{start}
	for len(queue) != 0 {
		// pp.Println("=== DEBUG START ======================================")

		curPos := queue[0]
		curProb := probs[curPos]
		queue = queue[1:]

		// pp.Println("curPos:", curPos, "curProb:", curProb)

		for _, path := range graph[curPos] {
			// pp.Println("test")
			// pp.Println("form:", curPos, "to:", path.to)
			nextPos := path.to
			nextProb := curProb * path.prob

			// 小さすぎるときもNG
			if nextProb < probs[end] {
				continue
			}

			// 更新したときだけ追加
			if nextProb > probs[nextPos] {
				// pp.Println("追加")
				// pp.Println("nextPos:", nextPos, "nextProb:", nextProb)
				probs[nextPos] = nextProb
				queue = append(queue, nextPos)
			}
		}
		// pp.Println("=== DEBUG END ========================================")
	}
	return probs[end]
}
