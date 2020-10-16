// 264. Ugly Number II
// https://leetcode.com/problems/ugly-number-ii/
package main

import (
	"github.com/k0kubun/pp"
)

// July LeetCoding Challenge Day 4
// https://leetcode.com/explore/featured/card/july-leetcoding-challenge/544/week-1-july-1st-july-7th/3380/
func main() {
	// pp.Println("=========================================")
	// pp.Println(nthUglyNumber(10))
	// pp.Println(12)
	// pp.Println("=========================================")
	// pp.Println(nthUglyNumber(12))
	// pp.Println(16)
	// pp.Println("=========================================")
	// pp.Println(nthUglyNumber(16))
	// pp.Println(25)
	// pp.Println("=========================================")
	// pp.Println(nthUglyNumber(21))
	// pp.Println(40)
	pp.Println("=========================================")
	pp.Println(nthUglyNumber(1689))
	pp.Println()
	pp.Println("=========================================")
}

func nthUglyNumber(n int) int {
	two, three, five := []int{1}, []int{1}, []int{1}
	for n > 1 {

	}
	return two[len(two)-1]
}

// func nthUglyNumber(n int) int {
// 	var num int
// 	for n != 0 {
// 		num++
// 		if isUglyNumber(num) {
// 			n--
// 		}
// 	}
// 	return num
// }

// func isUglyNumber(num int) bool {
// 	for num&1 == 0 {
// 		num >>= 1
// 	}
// 	for num%3 == 0 {
// 		num /= 3
// 	}
// 	for num%5 == 0 {
// 		num /= 5
// 	}
// 	return num == 1
// }
