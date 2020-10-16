// 392. Is Subsequence
// https://leetcode.com/problems/is-subsequence/
package main

import (
	"strings"

	"github.com/k0kubun/pp"
)

// June LeetCoding Challenge Day 9
// https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/540/week-2-june-8th-june-14th/3355/
func main() {
	pp.Println("=========================================")
	pp.Println(isSubsequence("abc", "ahbgdc"))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isSubsequence("axc", "ahbgdc"))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(isSubsequence("", "ahbgdc"))
	pp.Println(true)
	pp.Println("=========================================")
}

/*
 * 3. Pre-processing and binary search solution
 * - Time Complexity: O(N) for initialization, O(MlogN) for a query
 * - Space Complexity: O(N)
 *   where N: len(t), M: len(s)
 */
// func isSubsequence(s string, t string) bool {
// 	// initialization
// 	var indices [26][]int
// 	for i, c := range t {
// 		indices[c-'a'] = append(indices[c-'a'], i)
// 	}
//
// 	// query
// 	pointer := -1
// 	for _, c := range s {
// 		index := sort.Search(len(indices[c-'a']), func(i int) bool {
// 			return indices[c-'a'][i] > pointer
// 		})
// 		if index == len(indices[c-'a']) {
// 			return false
// 		}
// 		pointer = indices[c-'a'][index]
// 	}
// 	return true
// }

/*
 * 2. `strings.Index` solution
 * - Time Complexity: O(N+M)
 * - Space Complexity: O(1)
 *   where N: len(t), M: len(s)
 */
func isSubsequence(s string, t string) bool {
	for _, r := range s {
		index := strings.IndexRune(t, r)
		if index == -1 {
			return false
		}
		t = t[index+1:]
	}
	return true
}

/*
 * 1. Two pointer solution
 * - Time Complexity: O(N+M) (or O(N))
 * - Space Complexity: O(1)
 *   where N: len(t), M: len(s)
 */
// func isSubsequence(s string, t string) bool {
// 	var pointer int
// 	for i := range t {
// 		if pointer == len(s) {
// 			break
// 		}
// 		if s[pointer] == t[i] {
// 			pointer++
// 		}
// 	}
// 	return pointer == len(s)
// }
//
// func isSubsequence(s string, t string) bool {
// 	var pointer int
// 	for i := 0; pointer < len(s) && i < len(t); i++ {
// 		if t[i] == s[pointer] {
// 			pointer++
// 		}
// 	}
// 	return pointer == len(s)
// }
