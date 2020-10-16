// x. xxx
// https://leetcode.com/problems/xxx/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("=========================================")
	pp.Println(maxFreq("aababcaab", 2, 3, 4))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(maxFreq("aaaa", 1, 3, 3))
	pp.Println(2)
	pp.Println("=========================================")
	pp.Println(maxFreq("aabcabcab", 2, 2, 3))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(maxFreq("abcde", 2, 3, 3))
	pp.Println(0)
	pp.Println("=========================================")
}

// 1 <= s.length <= 10^5
// 1 <= maxLetters <= 26
// 1 <= minSize <= maxSize <= min(26, s.length)
// s only contains lowercase English letters.
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	// map[string]int{}
	// でカウントする
	// minSizeのみで十分

	// []int
	// と
	// map[int][]string
	// に変換する

	// string
	// ごとにmaxLettersを満たすか確認する

	subStringCounter := map[string]int{}
	for i := 0; i <= len(s)-minSize; i++ {
		subStringCounter[s[i:i+minSize]]++
	}

	subStrings := map[int][]string{}
	for str, count := range subStringCounter {
		subStrings[count] = append(subStrings[count], str)
	}

	keys := []int{}
	for key := range subStrings {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for i := len(keys) - 1; i >= 0; i-- {
		key := keys[i]
	LOOP:
		for _, subString := range subStrings[key] {
			m := map[rune]int{}
			count := 0
			for _, r := range subString {
				if m[r] == 0 {
					count++
					if count > maxLetters {
						continue LOOP
					}
					m[r]++
				}
			}
			// pp.Println("=== DEBUG START ======================================")
			// pp.Println(count)
			// pp.Println(subString)
			// pp.Println(key)
			// pp.Println("=== DEBUG END ========================================")
			return key
		}
	}

	return 0
}
