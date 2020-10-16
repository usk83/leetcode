// 121. Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	pp.Println(5)
	pp.Println("=========================================")
	pp.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	pp.Println(0)
	pp.Println("=========================================")
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	profit, min := 0, prices[0]
	for _, price := range prices[1:] {
		if price <= min {
			min = price
		} else if p := price - min; p > profit {
			profit = p
		}
	}
	return profit
}

/*
 * weekly_contest_174_20200201
 * An exercise before the contest
 * as 00.go
 */
// func maxProfit(prices []int) int {
// 	if len(prices) < 2 {
// 		return 0
// 	}
// 	lowest, profit := prices[0], 0
// 	for _, price := range prices[1:] {
// 		if price < lowest {
// 			lowest = price
// 		} else if p := price - lowest; p > profit {
// 			profit = p
// 		}
// 	}
// 	return profit
// }
