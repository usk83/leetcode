// 594. Longest Harmonious Subsequence
// https://leetcode.com/problems/longest-harmonious-subsequence/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	pp.Println(5)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

func findLHS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}

	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	var longest int
	for i := range keys[:len(keys)-1] {
		if keys[i]+1 == keys[i+1] {
			length := m[keys[i]] + m[keys[i+1]]
			if length > longest {
				longest = length
			}
		}
	}
	return longest
}

// func findLHS(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	sort.Ints(nums)
// 	longest, prevCount, cur, count := 0, 0, nums[0], 1
// 	for _, num := range nums[1:] {
// 		if num == cur {
// 			count++
// 		} else {
// 			if prevCount != 0 && count != 0 && prevCount+count > longest {
// 				longest = prevCount + count
// 			}
// 			if num-1 == cur {
// 				prevCount = count
// 			} else {
// 				prevCount = 0
// 			}
// 			cur, count = num, 1
// 		}
// 	}
// 	if prevCount != 0 && count != 0 && prevCount+count > longest {
// 		return prevCount + count
// 	}
// 	return longest
// }

// func findLHS(nums []int) int {
// 	memo := map[int][3]int{}
// 	var longest int

// 	for _, num := range nums {
// 		for i := num - 1; i <= num; i++ {
// 			m := memo[i]
// 			m[i-num+1] = 1
// 			m[2]++
// 			if m[0] == 1 && m[1] == 1 && m[2] > longest {
// 				longest = m[2]
// 			}
// 			memo[i] = m
// 		}
// 	}

// 	return longest
// }
