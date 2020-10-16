// 888. Fair Candy Swap
// https://leetcode.com/problems/fair-candy-swap/
// 13:30
// Runtime: 76 ms, faster than 54.17%
// Memory Usage: 6.8 MB, less than 100.00%
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	fmt.Println(fairCandySwap([]int{1, 1}, []int{2, 2}))
	fmt.Println([]int{1, 2})
	pp.Println("=========================================")
	fmt.Println(fairCandySwap([]int{1, 2}, []int{2, 3}))
	fmt.Println([]int{1, 2})
	pp.Println("=========================================")
	fmt.Println(fairCandySwap([]int{2}, []int{1, 3}))
	fmt.Println([]int{2, 3})
	pp.Println("=========================================")
	fmt.Println(fairCandySwap([]int{1, 2, 5}, []int{2, 4}))
	fmt.Println([]int{5, 4})
	pp.Println("=========================================")
}

/*
 * v3. Create only one set
 * > Runtime: 60 ms, faster than 100.00%
 * > Memory Usage: 6.7 MB, less than 100.00%
 */
// func fairCandySwap(a []int, b []int) []int {
// 	var sumDiff int
// 	for _, candy := range a {
// 		sumDiff += candy
// 	}
// 	bSet := make(map[int]struct{}, len(b))
// 	for _, candy := range b {
// 		bSet[candy] = struct{}{}
// 		sumDiff -= candy
// 	}
// 	swapDiff := sumDiff >> 1
// 	for _, candy := range a {
// 		if _, ok := bSet[candy-swapDiff]; ok {
// 			return []int{candy, candy - swapDiff}
// 		}
// 	}
// 	return nil
// }

/*
 * v2. Create two sets and iterate over shorter set
 */
func fairCandySwap(a []int, b []int) []int {
	var sumDiff int
	aSet, bSet := make(map[int]struct{}, len(a)), make(map[int]struct{}, len(b))
	for _, candy := range a {
		aSet[candy] = struct{}{}
		sumDiff += candy
	}
	for _, candy := range b {
		bSet[candy] = struct{}{}
		sumDiff -= candy
	}
	swapDiff := sumDiff >> 1
	for candy := range aSet {
		if _, ok := bSet[candy-swapDiff]; ok {
			return []int{candy, candy - swapDiff}
		}
	}
	return nil
}

/*
 * v1. Create two sets
 * > Runtime: 64 ms, faster than 91.67%
 * > Memory Usage: 6.7 MB, less than 100.00%
 */
// func fairCandySwap(a []int, b []int) []int {
// 	var sumDiff int
// 	aSet, bSet := make(map[int]struct{}, len(a)), make(map[int]struct{}, len(b))
// 	for _, candy := range a {
// 		aSet[candy] = struct{}{}
// 		sumDiff += candy
// 	}
// 	for _, candy := range b {
// 		bSet[candy] = struct{}{}
// 		sumDiff -= candy
// 	}
// 	swapDiff := sumDiff >> 1
// 	for candy := range aSet {
// 		if _, ok := bSet[candy-swapDiff]; ok {
// 			return []int{candy, candy - swapDiff}
// 		}
// 	}
// 	return nil
// }
