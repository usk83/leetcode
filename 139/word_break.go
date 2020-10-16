// 139. Word Break
// https://leetcode.com/problems/word-break/
package main

import (
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(wordBreak("leetcode", []string{"leet", "code"}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))
	pp.Println(false)
	pp.Println("=========================================")
}

/*
 * Optimized DP (early return)
 */
// func wordBreak(s string, wordDict []string) bool {
// 	reached := make([]bool, len(s)+1)
// 	reached[0] = true
// 	for i := range s {
// 		if !reached[i] {
// 			continue
// 		}
// 		for _, word := range wordDict {
// 			dist := i + len(word)
// 			if dist > len(s) || reached[dist] {
// 				continue
// 			}
// 			if strings.HasPrefix(s[i:], word) {
// 				if dist == len(s) {
// 					return true
// 				}
// 				reached[dist] = true
// 			}
// 		}
// 	}
// 	return false
// }

/*
 * DP
 */
// func wordBreak(s string, wordDict []string) bool {
// 	reached := make([]bool, len(s)+1)
// 	reached[0] = true
// 	for i := range s {
// 		if !reached[i] {
// 			continue
// 		}
// 		for _, word := range wordDict {
// 			if i+len(word) > len(s) || reached[i+len(word)] {
// 				continue
// 			}
// 			if strings.HasPrefix(s[i:], word) {
// 				reached[i+len(word)] = true
// 			}
// 		}
// 	}
// 	return reached[len(s)]
// }

/*
 * DFS - Backtracking with memoization
 */
func wordBreak(s string, wordDict []string) bool {
	failed := make([]bool, len(s)+1)
	var dfs func(index int) bool
	dfs = func(index int) bool {
		if index == len(s) {
			return true
		}
		if failed[index] {
			return false
		}
		for _, word := range wordDict {
			dist := index + len(word)
			if dist > len(s) || failed[dist] {
				continue
			}
			if strings.HasPrefix(s[index:], word) && dfs(dist) {
				return true
			}
		}
		failed[index] = true
		return false
	}
	return dfs(0)
}

/*
 * DFS - Backtracking (Time Limit Exceeded)
 */
// func wordBreak(s string, wordDict []string) bool {
// 	var dfs func(target string) bool
// 	dfs = func(target string) bool {
// 		if target == "" {
// 			return true
// 		}
// 		for _, word := range wordDict {
// 			if strings.HasPrefix(target, word) && dfs(target[len(word):]) {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// 	return dfs(s)
// }
