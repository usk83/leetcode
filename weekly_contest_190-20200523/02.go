// 5417. Maximum Number of Vowels in a Substring of Given Length
// https://leetcode.com/contest/weekly-contest-190/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
package main

import (
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxVowels("abciiidef", 3))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(maxVowels("aeiou", 2))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(maxVowels("leetcode", 3))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(maxVowels("rhythms", 4))
	pp.Println(0)
	pp.Println("=========================================")
	pp.Println(maxVowels("tryhard", 4))
	pp.Println(1)
	pp.Println("=========================================")
}

// 1 <= s.length <= 10^5
// s consists of lowercase English letters.
// 1 <= k <= s.length
func maxVowels(s string, k int) int {
	var count int
	for i := 0; i < k; i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	maxCount := count
	for start, i := 0, k; i < len(s); start, i = start+1, i+1 {
		switch s[start] {
		case 'a', 'e', 'i', 'o', 'u':
			count--
		}

		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}

		maxCount = max(maxCount, count)
	}
	return maxCount
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
