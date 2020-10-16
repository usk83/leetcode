// 5524. Maximum Profit of Operating a Centennial Wheel
// https://leetcode.com/contest/weekly-contest-208/problems/maximum-profit-of-operating-a-centennial-wheel/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(minOperationsMaxProfit([]int{8, 3}, 5, 6))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(minOperationsMaxProfit([]int{10, 9, 6}, 6, 4))
	pp.Println(7)
	pp.Println("=========================================")
	pp.Println(minOperationsMaxProfit([]int{3, 4, 0, 5, 1}, 1, 92))
	pp.Println(-1)
	pp.Println("=========================================")
	pp.Println(minOperationsMaxProfit([]int{10, 10, 6, 4, 7}, 3, 8))
	pp.Println(9)
	pp.Println("=========================================")
	pp.Println(minOperationsMaxProfit([]int{10, 10, 1, 0, 0}, 4, 4))
	pp.Println(5)
	pp.Println("=========================================")
	pp.Println(minOperationsMaxProfit([]int{0, 43, 37, 9, 23, 35, 18, 7, 45, 3, 8, 24, 1, 6, 37, 2, 38, 15, 1, 14, 39, 27, 4, 25, 27, 33, 43, 8, 44, 30, 38, 40, 20, 5, 17, 27, 43, 11, 6, 2, 30, 49, 30, 25, 32, 3, 18, 23, 45, 43, 30, 14, 41, 17, 42, 42, 44, 38, 18, 26, 32, 48, 37, 5, 37, 21, 2, 9, 48, 48, 40, 45, 25, 30, 49, 41, 4, 48, 40, 29, 23, 17, 7, 5, 44, 23, 43, 9, 35, 26, 44, 3, 26, 16, 31, 11, 9, 4, 28, 49, 43, 39, 9, 39, 37, 7, 6, 7, 16, 1, 30, 2, 4, 43, 23, 16, 39, 5, 30, 23, 39, 29, 31, 26, 35, 15, 5, 11, 45, 44, 45, 43, 4, 24, 40, 7, 36, 10, 10, 18, 6, 20, 13, 11, 20, 3, 32, 49, 34, 41, 13, 11, 3, 13, 0, 13, 44, 48, 43, 23, 12, 23, 2}, 43, 54))
	pp.Println(993)
	pp.Println("=========================================")
}

// n == customers.length
// 1 <= n <= 105
// 0 <= customers[i] <= 50
// 1 <= boardingCost, runningCost <= 100
func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	onePref := boardingCost<<2 - runningCost
	if onePref <= 0 {
		return -1
	}
	// 利益最大値を保持しておく
	var maxProf, minOp, prof, op, waiting int
	// 都度都度乗せられるだけ乗せる
	// オペレーション回数と利益を計算する
	process := func() {
		op++
		if waiting >= 4 {
			prof += onePref
			waiting -= 4
		} else {
			prof += boardingCost*waiting - runningCost
			waiting = 0
		}
		if prof > maxProf {
			maxProf = prof
			minOp = op
		}
	}
	for _, c := range customers {
		waiting += c
		process()
	}
	// 最後残った人を処理する
	for waiting != 0 {
		process()
	}
	return minOp
}
