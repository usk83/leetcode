// 567. Permutation in String
// https://leetcode.com/problems/permutation-in-string/
package main

import (
	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 18
// https://leetcode.com/explore/featured/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3333/
func main() {
	pp.Println("=========================================")
	pp.Println(checkInclusion("ab", "eidbaooo"))
	pp.Println(true)
	pp.Println("=========================================")
	pp.Println(checkInclusion("ab", "eidboaoo"))
	pp.Println(false)
	pp.Println("=========================================")
}

func checkInclusion(s1 string, s2 string) bool {
	count := len(s1)
	var counts [26]int
	for _, r := range s1 {
		counts[r-'a']++
	}

	var left, right int
	for left < len(s2)-len(s1)+1 && right < len(s2) {
		if counts[s2[right]-'a'] > 0 {
			counts[s2[right]-'a']--
			count--
			right++
		} else {
			counts[s2[left]-'a']++
			count++
			left++
		}
		if count == 0 {
			return true
		}
	}
	return false
}
