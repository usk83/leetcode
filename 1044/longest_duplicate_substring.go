// https://leetcode.com/problems/longest-duplicate-substring/
// 1044. Longest Duplicate Substring
package main

import (
	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 19
// https://leetcode.com/explore/featured/card/june-leetcoding-challenge/541/week-3-june-15th-june-21st/3365/
func main() {
	pp.Println("=========================================")
	pp.Println(longestDupSubstring("banana"))
	pp.Println("ana")
	pp.Println("=========================================")
	pp.Println(longestDupSubstring("abcd"))
	pp.Println("")
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

// 2 <= S.length <= 10^5
// S consists of lowercase English letters.

func longestDupSubstring(s string) string {
	l, r := 1, len(s)
	pos := make([]int, len(s))
LOOP:
	for l < r {
		m := (l + r) >> 1
		// fmt.Printf("left: %d, right: %d, middle: %d\n", l, r, m)
		set := map[string]struct{}{}
		for i := range s[:len(s)-m+1] {
			current := s[i : i+m]
			// pp.Println("i:", i)
			// pp.Println("current:", current)
			if _, ok := set[current]; ok {
				pos[m] = i
				l = m + 1
				continue LOOP
			}
			set[current] = struct{}{}
		}
		r = m
	}

	// fmt.Printf("left: %d, right: %d\n", l, r)
	start := pos[l-1]
	end := start + l - 1
	// fmt.Printf("start: %d, end: %d\n", start, end)
	return s[start:end]
}

/**
 * 0. DP - TLE
 */
// func longestDupSubstring(s string) string {
// 	positions := [26][]int{}
// 	for i, r := range s {
// 		positions[r-'a'] = append(positions[r-'a'], i)
// 	}
//
// 	dp := [2][]int{
// 		make([]int, len(s)),
// 		make([]int, len(s)),
// 	}
// 	var longestLength, endIndex int
// 	for i, r := range s {
// 		for j := i + 1; j < len(s); j++ {
// 			dp[i&1][j] = 0
// 		}
// 		var outdated int
// 		for _, pos := range positions[r-'a'] {
// 			if i >= pos {
// 				outdated++
// 				continue
// 			}
// 			dp[i&1][pos] = dp[i&1^1][pos-1] + 1
// 			if dp[i&1][pos] > longestLength {
// 				longestLength = dp[i&1][pos]
// 				endIndex = pos
// 			}
// 		}
// 		positions[r-'a'] = positions[r-'a'][outdated:]
// 	}
// 	return s[endIndex-longestLength+1 : endIndex+1]
// }
