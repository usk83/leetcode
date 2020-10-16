// 438. Find All Anagrams in a String
// https://leetcode.com/problems/find-all-anagrams-in-a-string/
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 17
// https://leetcode.com/explore/featured/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3332/
func main() {
	pp.Println("=========================================")
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println([]int{0, 6})
	pp.Println("=========================================")
	fmt.Println(findAnagrams("abab", "ab"))
	fmt.Println([]int{0, 1, 2})
	pp.Println("=========================================")
}

func findAnagrams(s string, p string) []int {
	// 事前処理
	// 	[26]int, int
	// right が一致するとき right を動かす
	// 	減らす
	// right が一致しないとき left を動かす
	// 	増やす
	// all zeroになったらそのときの left を答えに追加
	// leftかrightが最後の可能性を超えたら終了

	count := len(p)
	var counts [26]int
	for _, r := range p {
		counts[r-'a']++
	}

	var indices []int
	var left, right int
	for left < len(s)-len(p)+1 && right < len(s) {
		if counts[s[right]-'a'] > 0 {
			counts[s[right]-'a']--
			count--
			right++
		} else {
			counts[s[left]-'a']++
			count++
			left++
		}
		if count == 0 {
			indices = append(indices, left)
		}
	}
	return indices
}
