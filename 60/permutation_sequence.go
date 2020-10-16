// 60. Permutation Sequence
// https://leetcode.com/problems/permutation-sequence/
package main

import (
	"strconv"

	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 20
// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/541/week-3-june-15th-june-21st/3366/
func main() {
	pp.Println("=========================================")
	pp.Println(getPermutation(3, 3))
	pp.Println("213")
	pp.Println("=========================================")
	pp.Println(getPermutation(4, 9))
	pp.Println("2314")
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

var factorials = [9]int{0, 1, 2, 6, 24, 120, 720, 5040, 40320}

func getPermutation(n int, k int) string {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}

	k -= 1

	var permutation int
	for i := n - 1; i > 0; i-- {
		permutation *= 10
		j := k / factorials[i]
		k %= factorials[i]
		permutation += nums[j]
		for k := j - 1; k >= 0; k-- {
			nums[k+1] = nums[k]
		}
		nums = nums[1:]
	}
	permutation = permutation*10 + nums[0]
	return strconv.Itoa(permutation)
}
