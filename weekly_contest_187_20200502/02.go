// 5401. Check If All 1's Are at Least Length K Places Away
// https://leetcode.com/contest/weekly-contest-187/problems/check-if-all-1s-are-at-least-length-k-places-away/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
	pp.Println()
	pp.Println()
	pp.Println("=========================================")
}

func kLengthApart(nums []int, k int) bool {
	prev := -1
	for i, num := range nums {
		if num == 1 {
			if prev != -1 && i-prev-1 < k {
				return false
			}
			prev = i
		}
	}
	return true
}
