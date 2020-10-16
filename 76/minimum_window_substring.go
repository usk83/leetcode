// 76. Minimum Window Substring
// https://leetcode.com/problems/minimum-window-substring/
package main

import (
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println(minWindow("ADOBECODEBANC", "ABC"))
	pp.Println("Expected:", "BANC")

	pp.Println(minWindow("a", "a"))
	pp.Println("Expected:", "a")

	pp.Println(minWindow("a", "aa"))
	pp.Println("Expected:", "")

	pp.Println(minWindow("bbaac", "aba"))
	pp.Println("Expected:", "baa")

	pp.Println(minWindow("aaaaaaaaaaaabbbbbcdd", "abcdd"))
	pp.Println("Expected:", "abbbbbcdd")

	pp.Println(minWindow("ask_not_what_your_country_can_do_for_you_ask_what_you_can_do_for_your_country", "ask_country"))
	pp.Println("Expected:", "sk_not_what_your_c")
}

func minWindow(s string, t string) string {
	dp := [2]map[string]int{}
	dp[0] = map[string]int{}
	minIndex := 0
	minLength := 0
	for i := len(s) - 1; i >= 0; i-- {
		currentChar := string(s[i])
		if strings.Index(t, currentChar) < 0 {
			continue
		}

		dp[0][t] = i
		dp[1] = map[string]int{}

		for key, endIndex := range dp[0] {
			var newKey string
			if index := strings.Index(key, currentChar); index < 0 {
				newKey = key
			} else {
				newKey = key[:index] + key[index+1:]
			}
			if newKey != "" {
				index, found := dp[1][newKey]
				if !found || endIndex < index {
					dp[1][newKey] = endIndex
				}
				continue
			}
			length := endIndex - i + 1
			if minLength == 0 || length < minLength {
				minIndex = i
				minLength = length
			}
		}

		dp[0] = dp[1]
	}

	return s[minIndex : minIndex+minLength]
}

func minWindowV1(s string, t string) string {
	var target int64 = 0
	for i := len(t) - 1; i >= 0; i-- {
		target = target | 1<<(t[i]-65)
	}

	dp := map[int64]int{}
	keys := make([]int64, 52)
	keysLength := 0
	minIndex := 0
	minLength := 0
	for i := len(s) - 1; i >= 0; i-- {
		var current int64 = 1 << (s[i] - 65)
		if target&current == 0 {
			continue
		}
		if current == target {
			return string(s[i])
		}

		for k := range dp {
			if k&current != 0 {
				continue
			}
			keys[keysLength] = k
			keysLength++
		}

		dp[current] = i

		for j := 0; j < keysLength; j++ {
			endIndex := dp[keys[j]]
			delete(dp, keys[j])
			if newKey := keys[j] | current; newKey != target {
				dp[newKey] = endIndex
				continue
			}
			length := endIndex - i + 1
			if minLength == 0 || length < minLength {
				minIndex = i
				minLength = length
			}
		}

		keysLength = 0
	}

	return s[minIndex : minIndex+minLength]
}
