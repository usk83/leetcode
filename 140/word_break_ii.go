// 140. Word Break II
// https://leetcode.com/problems/word-break-ii/
package main

import (
	"fmt"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	pp.Println([]string{
		"cats and dog",
		"cat sand dog",
	})
	pp.Println("=========================================")
	pp.Println(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
	pp.Println([]string{
		"pine apple pen apple",
		"pineapple pen apple",
		"pine applepen apple",
	})
	pp.Println("=========================================")
	pp.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	pp.Println([]string{})
	pp.Println("=========================================")
}

func wordBreak(s string, wordDict []string) []string {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = []int{}
	}
	dp[0] = []int{0}

	for i := range s {
		if len(dp[i]) == 0 {
			continue
		}
		for _, word := range wordDict {
			if i+len(word) > len(s) {
				continue
			}
			if strings.HasPrefix(s[i:], word) {
				dp[i+len(word)] = append(dp[i+len(word)], i)
			}
		}
	}

	var dfs func(i int) []string
	dfs = func(i int) []string {
		words := []string{}
		for _, j := range dp[i] {
			if j == 0 {
				words = append(words, s[:i])
				continue
			}
			for _, word := range dfs(j) {
				words = append(words, fmt.Sprintf("%s %s", word, s[j:i]))
			}
		}
		return words
	}

	return dfs(len(s))
}
