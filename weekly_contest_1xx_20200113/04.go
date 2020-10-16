// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println(minTaps(5, []int{3, 0, 1, 1, 1, 0}))
	// pp.Println("=========================================")
	// pp.Println(minTaps(5, []int{3, 4, 1, 1, 0, 0}))
	// pp.Println(1)
	// pp.Println("=========================================")
	// pp.Println(minTaps(3, []int{0, 0, 0, 0}))
	// pp.Println(-1)
	// pp.Println("=========================================")
	// pp.Println(minTaps(7, []int{1, 2, 1, 0, 2, 1, 0, 1}))
	// pp.Println(3)
	// pp.Println("=========================================")
	// pp.Println(minTaps(8, []int{4, 0, 0, 0, 0, 0, 0, 0, 4}))
	// pp.Println(2)
	// pp.Println("=========================================")
	// pp.Println(minTaps(8, []int{4, 0, 0, 0, 4, 0, 0, 0, 4}))
	// pp.Println(1)
	// pp.Println("=========================================")
}

// 1 <= n <= 10^4
// ranges.length == n + 1
// 0 <= ranges[i] <= 100
func minTaps(n int, ranges []int) int {

	// 自分でしか到達できないやつはマスト
	// 自分の領域をカバーするタップのリストを作成する

	// マストのやつを考慮したときにまだカバーされていないリストを考える

	// 下位互換を消す

	// マストのやつがあるときはそれ以外のやつを消して良い
	// その中でどれを取るのが一番効率いいか

	type data struct {
		taps map[int]bool
	}

	garden := map[int]data{}
	for i := n - 1; i >= 0; i-- {
		garden[i] = data{
			taps: map[int]bool{},
		}
	}
	area := make([][2]int, len(ranges))

	for i, covers := range ranges {
		start, end := i-covers, i+covers
		if start < 0 {
			start = 0
		}
		if end > n {
			end = n
		}

		area[i][0] = start
		area[i][1] = end
		for j := start; j < end; j++ {
			garden[j].taps[i] = true
		}
	}

	var opened int

	nes := []int{}
	for _, g := range garden {
		if l := len(g.taps); l == 0 {
			return -1
		} else if l == 1 {
			var tap int
			for t := range g.taps {
				tap = t
			}
			nes = append(nes, tap)
		}
	}

	// fmt.Println(nes)

	for _, n := range nes {
		opened++
		for _, cover := range area[n] {
			delete(garden, cover)
		}
	}

	// fmt.Println(garden)

	return opened
}
