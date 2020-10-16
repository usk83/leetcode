// 392. Is Subsequence
// https://leetcode.com/problems/is-subsequence/
package main

import (
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(isSubsequence("abc", "ahbgdc"))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(isSubsequence("axc", "ahbgdc"))
	pp.Println(false)
	pp.Println("=========================================")
}

// func isSubsequence(s string, t string) bool {
// 	sLen := len(s)
// 	if sLen == 0 {
// 		return true
// 	}
// 	tLen := len(t)
// 	if tLen == 0 {
// 		return false
// 	}
// 	dp := make([]int, tLen)

// 	if s[0] == t[0] {
// 		dp[0] = 1
// 		s = s[1:]
// 	} else {
// 		dp[0] = 0
// 	}

// 	for i := 1; i < tLen; i++ {
// 		if t[i] == s[0] {
// 			dp[i] = dp[i-1] + 1
// 			s = s[1:]
// 		} else {
// 			dp[i] = dp[i-1]
// 		}
// 		if dp[i] == sLen {
// 			return true
// 		}
// 	}

// 	return false
// }

// func isSubsequence(s string, t string) bool {
// 	sLen, currentIndex := len(s), 0
// 	for i := 0; i < len(t); i++ {
// 		if currentIndex == sLen {
// 			break
// 		}
// 		if t[i] == s[currentIndex] {
// 			currentIndex++
// 		}
// 	}
// 	return currentIndex == sLen
// }

// func isSubsequence(s string, t string) bool {
// 	sLen := len(s)
// 	if len(s) == 0 {
// 		return true
// 	}
// 	tLen := len(t)
// 	if tLen == 0 {
// 		return false
// 	}
// 	dp := make([]int, tLen)

// 	// dp[0] = s[0] == t[0] ? 1 : 0
// 	if s[0] == t[0] {
// 		dp[0] = 1
// 	} else {
// 		dp[0] = 0
// 	}

// 	for i := 1; i < tLen; i++ {
// 		if t[i] == s[dp[i-1]] {
// 			dp[i] = dp[i-1] + 1
// 		} else {
// 			dp[i] = dp[i-1]
// 		}
// 		if dp[i] == sLen {
// 			return true
// 		}
// 	}

// 	return false
// }

func isSubsequence(s string, t string) bool {
	for _, c := range s {
		index := strings.Index(t, string(c))
		if index == -1 {
			return false
		}
		t = t[index+1:]
	}
	return true
}

// // isSubsequence using recursion
// func isSubsequence(s string, t string) bool {
// 	if s == "" {
// 		return true
// 	}
// 	i := strings.Index(t, string(s[0]))
// 	if i == -1 {
// 		return false
// 	}
// 	return isSubsequence(s[1:], t[i+1:])
// }
