// 72. Edit Distance
// https://leetcode.com/problems/edit-distance/
package main

import (
	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 31
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/538/week-5-may-29th-may-31st/3346/
func main() {
	pp.Println("=========================================")
	pp.Println(minDistance("horse", "ros"))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(minDistance("intention", "execution"))
	pp.Println(5)
	pp.Println("=========================================")
	pp.Println(minDistance("dinitrophenylhydrazine", "acetylphenylhydrazine"))
	pp.Println(6)
	pp.Println("=========================================")
	pp.Println(minDistance("teacher", "tepache"))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(minDistance("acher", "pache"))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(minDistance("pneumonoultramicroscopicsilicovolcanoconiosis", "pneumonoconiosis"))
	pp.Println(29)
	pp.Println("=========================================")
}

/*
 * 1. The first solution
 */
func minDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}

	if len(word1) < len(word2) {
		word1, word2 = word2, word1
	}
	var commonPrefixLength int
	for i := 0; i < len(word2); i++ {
		if word1[i] != word2[i] {
			break
		}
		commonPrefixLength++
	}

	var commonSuffixLength int
	for i := 0; i < len(word2); i++ {
		if word1[len(word1)-1-i] != word2[len(word2)-1-i] {
			break
		}
		commonSuffixLength++
	}

	if commonPrefixLength+commonSuffixLength >= len(word2) {
		return len(word1) - len(word2)
	}

	dp := make([][]int, len(word1)+1-commonPrefixLength-commonSuffixLength)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1-commonPrefixLength-commonSuffixLength)
		dp[i][0] = i
	}

	for i := range dp[0][1:] {
		dp[0][i+1] = i + 1
	}

	for i := range dp[1:] {
		for j := range dp[i+1][1:] {
			if word1[i+commonPrefixLength] == word2[j+commonPrefixLength] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min3(dp[i][j], dp[i+1][j], dp[i][j+1]) + 1
			}
		}
	}

	return dp[len(dp)-1][len(dp[len(dp)-1])-1]
}

func min3(x, y, z int) int {
	if x < y {
		if x < z {
			return x
		}
		return z
	}
	if y < z {
		return y
	}
	return z
}
