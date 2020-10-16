package main

import "github.com/k0kubun/pp"

func main() {
	pp.Println(minCostToMoveChips([]int{1, 2, 3}))
	pp.Println(1)

	pp.Println(minCostToMoveChips([]int{2, 2, 2, 3, 3}))
	pp.Println(2)
}

// 1 <= chips.length <= 100
// 1 <= chips[i] <= 10^9

func minCostToMoveChips(chips []int) int {
	// 偶数と奇数の数を数える
	// それぞれはコスト0でいける
	// 小さい方の数字が答え

	evens := []int{}
	odds := []int{}
	for _, c := range chips {
		if c%2 == 0 {
			evens = append(evens, c)
		} else {
			odds = append(odds, c)
		}
	}

	if len(evens) < len(odds) {
		return len(evens)
	}

	return len(odds)
}
