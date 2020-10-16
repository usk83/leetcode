// 535?. Design a Stack With Increment Operation
// https://leetcode.com/contest/weekly-contest-180/problems/maximum-performance-of-a-team/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	// pp.Println("=========================================")
	// pp.Println(maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 2))
	// pp.Println(60)
	// pp.Println("=========================================")
	// pp.Println(maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 3))
	// pp.Println(68)
	// pp.Println("=========================================")
	// pp.Println(maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 4))
	// pp.Println(72)
	pp.Println("=========================================")
	pp.Println(maxPerformance(3, []int{2, 8, 2}, []int{2, 7, 1}, 2))
	pp.Println(56)
	pp.Println("=========================================")
}

// 1 <= n <= 10^5
// speed.length == n
// efficiency.length == n
// 1 <= speed[i] <= 10^5
// 1 <= efficiency[i] <= 10^8
// 1 <= k <= n
func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	const BIG_NUM = 1000000000 + 7

	// efficiencyを下げないなら、speed最大のやつ
	// そいつよりsppedが大きくて、efficiencyを下げるやつ

	// 最初のやつに対して、efficiencyを下げない場合は選択する
	// 下げるとき
	// dp := make([])
	// for i := 0; i < n; i++ {
	// 	for j := i; j < n; j++ {
	// 	}
	// }

	// speed順にソートする
	// k人取る
	// efficiencyを上げる場合、
	// 最初のやつと比較する

	type engineer struct {
		speed      int
		efficiency int
	}

	engineers := make([]engineer, n)
	for i := 0; i < n; i++ {
		engineers[i] = engineer{speed[i], efficiency[i]}
	}

	// speedソート
	sort.Slice(engineers, func(i, j int) bool {
		return engineers[i].speed > engineers[j].speed
	})

	// 本当はPriority Queue
	selected := engineers[:k]
	speedSum := 0
	for _, e := range selected {
		speedSum += e.speed
	}
	engineers = engineers[k:]

	sort.Slice(selected, func(i, j int) bool {
		return selected[i].efficiency < selected[j].efficiency
	})

	for _, eng := range engineers {
		// fmt.Println(selected)

		if eng.efficiency <= selected[0].efficiency {
			continue
		}

		curDiv := speedSum * selected[0].efficiency / BIG_NUM
		curMod := speedSum * selected[0].efficiency % BIG_NUM

		nextMin := eng.efficiency
		if len(selected) > 1 {
			nextMin = min(nextMin, selected[1].efficiency)
		}

		nextDiv := (speedSum - selected[0].speed + eng.speed) * nextMin / BIG_NUM
		nextMod := (speedSum - selected[0].speed + eng.speed) * nextMin % BIG_NUM

		// pp.Println(eng.speed, eng.efficiency)
		// pp.Println(speedSum)
		// pp.Println(selected[0].efficiency)
		// pp.Println(speedSum - selected[0].speed + eng.speed)
		// pp.Println(nextMin)

		if nextDiv < curDiv || nextMod < curMod {
			continue
		}

		speedSum = speedSum - selected[0].speed + eng.speed
		selected[0] = eng
		sort.Slice(selected, func(i, j int) bool {
			return selected[i].efficiency < selected[j].efficiency
		})
	}
	// fmt.Println(selected)

	return (speedSum * selected[0].efficiency) % BIG_NUM
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
