// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println(minTaps(5, []int{1, 1, 1, 1, 1, 1}))
	// pp.Println("=========================================")
	// pp.Println(minTaps(5, []int{3, 4, 1, 1, 0, 0}))
	// // pp.Println(1)
	// pp.Println("=========================================")
	// pp.Println(minTaps(3, []int{0, 0, 0, 0}))
	// // pp.Println(-1)
	// pp.Println("=========================================")
	// pp.Println(minTaps(7, []int{1, 2, 1, 0, 2, 1, 0, 1}))
	// // pp.Println(3)
	// pp.Println("=========================================")
	// pp.Println(minTaps(8, []int{4, 0, 0, 0, 0, 0, 0, 0, 4}))
	// // pp.Println(2)
	// pp.Println("=========================================")
	// pp.Println(minTaps(8, []int{4, 0, 0, 0, 4, 0, 0, 0, 4}))
	// // pp.Println(1)
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
		if covers == 0 {
			continue
		}

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

	// fmt.Println(garden)

	// 下位互換を削除
	unnes := []int{}
	for i, cover := range area {
		if cover[0] == 0 && cover[1] == 0 {
			continue
		}
		// 自分のカバー領域に自分よりもrangeが広いやつがいればunnes
		starts := garden[cover[0]].taps
		ends := garden[cover[1]-1].taps

		for s := range starts {
			if ends[s] && s != i {
				unnes = append(unnes, i)
				break
			}
		}
	}

	// fmt.Println(unnes)

	var opened int

	// nes := []int{}
	nes := map[int]bool{}
	for _, g := range garden {
		for _, un := range unnes {
			delete(g.taps, un)
		}

		if l := len(g.taps); l == 0 {
			return -1
		} else if l == 1 {
			var tap int
			for t := range g.taps {
				tap = t
			}
			// nes = append(nes, tap)
			nes[tap] = true
		}
	}

	// fmt.Println(nes)

	fmt.Println(garden)

	for n := range nes {
		opened++
		pp.Println(area[n][0], area[n][1])
		for i := area[n][0]; i < area[n][1]; i++ {
			delete(garden, i)
		}
	}

	fmt.Println(garden)

	return opened
}
