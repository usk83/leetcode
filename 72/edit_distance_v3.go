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
 * 3. Optimized solution
 */
func minDistance(word1 string, word2 string) int {
	if len(word1) < len(word2) {
		word1, word2 = word2, word1
	}

	length1, length2 := len(word1), len(word2)

	var commonPrefixLength int
	for i := 0; i < length2; i++ {
		if word1[i] != word2[i] {
			break
		}
		commonPrefixLength++
	}

	var commonSuffixLength int
	for i := 0; i < length2-commonPrefixLength; i++ {
		if word1[length1-1-i] != word2[length2-1-i] {
			break
		}
		commonSuffixLength++
	}

	if commonPrefixLength+commonSuffixLength == length2 {
		return length1 - length2
	}

	length1, length2 = length1-commonPrefixLength-commonSuffixLength, length2-commonPrefixLength-commonSuffixLength

	dp := [][]int{
		make([]int, length2+1),
		make([]int, length2+1),
	}

	for i := range dp[0] {
		dp[0][i] = i
	}

	for i := 0; i < length1; i++ {
		dp[i&1^1][0] = i + 1
		for j := 0; j < length2; j++ {
			if word1[commonPrefixLength+i] == word2[commonPrefixLength+j] {
				dp[i&1^1][j+1] = dp[i&1][j]
			} else {
				dp[i&1^1][j+1] = min3(dp[i&1][j], dp[i&1^1][j], dp[i&1][j+1]) + 1
			}
		}
	}

	return dp[length1&1][length2]
}

func min3(x, y, z int) int {
	if y < x {
		x = y
	}
	if z < x {
		x = z
	}
	return x
}
