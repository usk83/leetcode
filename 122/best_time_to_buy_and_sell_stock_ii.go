// 122. Best Time to Buy and Sell Stock II
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	pp.Println(7)
	pp.Println("=========================================")
	pp.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	pp.Println(4)
	pp.Println("=========================================")
	pp.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	pp.Println(0)
	pp.Println("=========================================")
}

/*
 * Shorten greedy solution
 */
func maxProfit(prices []int) int {
	var profit int
	for i := len(prices) - 1; i > 0; i-- {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

/*
 * DP solution if it cannot find greedy way
 */
// func maxProfit(prices []int) int {
// 	dp := map[int]int{
// 		-1: 0, // 買わない
// 	} // key: holding price, value: profit
// 	for _, current := range prices {
// 		next := map[int]int{}
// 		best := map[int]int{} // key: profit, value: holding price
//
// 		update := func(price, profit int) {
// 			if price == -1 {
// 				next[price] = max(next[price], profit)
// 				return
// 			}
//
// 			if provisional, ok := best[profit]; ok {
// 				if provisional < price {
// 					return
// 				}
// 				delete(next, provisional)
// 			}
// 			next[price] = max(next[price], profit)
// 			best[profit] = price
// 		}
//
// 		for holding, profit := range dp {
// 			switch {
// 			case holding == -1: // 持ってないとき
// 				update(current, profit) // 買う
// 				update(holding, profit) // 買わない（維持）
// 			case holding == current: // 同じとき
// 				update(holding, profit) // 維持
// 			case holding > current: // 下がったとき
// 				update(holding, profit) // 維持
// 			case holding < current: // 上がったとき
// 				update(-1, profit+current-holding) // 今売る
// 				update(holding, profit)            // まだ売らない（維持）
// 			}
// 		}
// 		dp = next
// 	}
//
// 	var maxProfit int
// 	for _, profit := range dp {
// 		maxProfit = max(maxProfit, profit)
// 	}
// 	return maxProfit
// }

/*
 * DP solution version 1 (TLE)
 */
// func maxProfit__(prices []int) int {
// 	dp := map[int]int{
// 		-1: 0, // 買わない
// 	} // key: holding price, value: profit
// 	for _, current := range prices {
// 		next := map[int]int{}
// 		for holding, profit := range dp {
// 			switch {
// 			case holding == -1: // 持ってないとき
// 				next[current] = max(next[current], profit) // 買う
// 				next[holding] = max(next[holding], profit) // 買わない（維持）
// 			case holding == current: // 同じとき
// 				next[holding] = max(next[holding], profit) // 維持
// 			case holding > current: // 下がったとき
// 				next[holding] = profit // 維持
// 			case holding < current: // 上がったとき
// 				next[-1] = max(next[-1], profit+current-holding) // 今売る
// 				next[holding] = profit                           // まだ売らない（維持）
// 			}
// 		}
// 		dp = next
// 	}
//
// 	var maxProfit int
// 	for _, profit := range dp {
// 		maxProfit = max(maxProfit, profit)
// 	}
// 	return maxProfit
// }

// helper for DP solutions
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

/*
 * The first solution
 */
// func maxProfit(prices []int) int {
// 	if len(prices) < 2 {
// 		return 0
// 	}
// 	profit, current := 0, prices[0]
// 	for _, price := range prices[1:] {
// 		if price > current {
// 			profit += price - current
// 		}
// 		current = price
// 	}
// 	return profit
// }
