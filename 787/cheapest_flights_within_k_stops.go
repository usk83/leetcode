// 787. Cheapest Flights Within K Stops
// https://leetcode.com/problems/cheapest-flights-within-k-stops/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 14
// https://leetcode.com/explore/featured/card/june-leetcoding-challenge/540/week-2-june-8th-june-14th/3360/
func main() {
	pp.Println("=========================================")
	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1))
	fmt.Println(200)
	pp.Println("=========================================")
	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0))
	fmt.Println(500)
	pp.Println("=========================================")
	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 2}, {1, 2, 1}, {2, 0, 10}}, 1, 2, 1))
	fmt.Println(1)
	pp.Println("=========================================")
	fmt.Println(findCheapestPrice(4, [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}, 0, 3, 1))
	fmt.Println(6)
	pp.Println("=========================================")
}

// The number of nodes n will be in range [1, 100], with nodes labeled from 0 to n - 1.
// The size of flights will be in range [0, n * (n - 1) / 2].
// The format of each flight will be (src, dst, price).
// The price of each flight will be in the range [1, 10000].
// k is in the range of [0, n - 1].
// There will not be any duplicated flights or self cycles.

/**
 * 1. The Shortest Path Faster Algorithm (SPFA)
 */
func findCheapestPrice(n int, rawFlights [][]int, src int, dst int, k int) int {
	if src == dst {
		return 0
	}
	flights := make([][][2]int, n)
	for _, rawFlight := range rawFlights {
		if rawFlight[1] != src {
			flights[rawFlight[0]] = append(flights[rawFlight[0]], [2]int{rawFlight[1], rawFlight[2]})
		}
	}
	cheapestPrices, queue := make([]int, n), [][3]int{{src, 0, 0}}
	for len(queue) != 0 {
		for _, flight := range flights[queue[0][0]] {
			price := queue[0][1] + flight[1]
			if cheapestPrices[dst] != 0 && price >= cheapestPrices[dst] {
				continue
			}
			if cheapestPrices[flight[0]] == 0 || price < cheapestPrices[flight[0]] {
				cheapestPrices[flight[0]] = price
				if stopCount := queue[0][2] + 1; stopCount <= k {
					queue = append(queue, [3]int{flight[0], price, stopCount})
				}
			}
		}
		queue = queue[1:]
	}
	if cheapestPrices[dst] == 0 {
		return -1
	}
	return cheapestPrices[dst]
}
