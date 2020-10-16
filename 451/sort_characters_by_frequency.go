// 451. Sort Characters By Frequency
// https://leetcode.com/problems/sort-characters-by-frequency/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 22
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/537/week-4-may-22nd-may-28th/3337/
func main() {
	pp.Println("=========================================")
	pp.Println(frequencySort("tree"))
	pp.Println("eert")
	pp.Println("=========================================")
	pp.Println(frequencySort("cccaaa"))
	pp.Println("cccaaa")
	pp.Println("=========================================")
	pp.Println(frequencySort("Aabb"))
	pp.Println("bbAa")
	pp.Println("=========================================")
	pp.Println(frequencySort("2a554442f544asfasssffffasss"))
	pp.Println()
	pp.Println("=========================================")
}

/*
 * 4. Using an slice of structs
 */
type charCount struct {
	char  byte
	count int
}

func frequencySort(s string) string {
	var charCounts [128]charCount
	for i := range s {
		charCounts[s[i]].char = s[i]
		charCounts[s[i]].count++
	}
	sort.Slice(charCounts[:], func(i, j int) bool {
		return charCounts[i].count > charCounts[j].count
	})
	result := make([]byte, 0, len(s))
	for _, cc := range charCounts {
		// if cc.count == 0 { // optimisation
		// 	break
		// }
		for ; cc.count > 0; cc.count-- {
			result = append(result, cc.char)
		}
	}
	return string(result)
}

/*
 * 3. Using fixed size array instead of map
 */
// func frequencySort(s string) string {
// 	var runeFreq [128]int
// 	for _, r := range s {
// 		runeFreq[r]++
// 	}
// 	freqCount := map[int][]byte{}
// 	var freqCountKeys []int
// 	for code, c := range runeFreq {
// 		if freqCount[c] == nil {
// 			freqCountKeys = append(freqCountKeys, c)
// 		}
// 		freqCount[c] = append(freqCount[c], byte(code))
// 	}
// 	sort.Sort(sort.Reverse(sort.IntSlice(freqCountKeys)))
// 	result := make([]byte, 0, len(s))
// 	for _, count := range freqCountKeys {
// 		for _, r := range freqCount[count] {
// 			for i := 0; i < count; i++ {
// 				result = append(result, r)
// 			}
// 		}
// 	}
// 	return string(result)
// }

/*
 * 2. Sort by an external factor
 * - slow
 */
// func frequencySort(s string) string {
// 	bytes := []byte(s)
// 	freq := map[byte]int{}
// 	for _, b := range bytes {
// 		freq[b]++
// 	}
// 	sort.Slice(bytes, func(i, j int) bool {
// 		if freq[bytes[i]] == freq[bytes[j]] {
// 			return bytes[i] < bytes[j]
// 		}
// 		return freq[bytes[i]] > freq[bytes[j]]
// 	})
// 	return string(bytes)
// }

/*
 * 1. The first solution
 */
// func frequencySort(s string) string {
// 	freqRune := map[rune]int{}
// 	for _, r := range s {
// 		freqRune[r]++
// 	}
// 	freqCount := map[int][]rune{}
// 	var freqCountKeys []int
// 	for r, c := range freqRune {
// 		if freqCount[c] == nil {
// 			freqCountKeys = append(freqCountKeys, c)
// 		}
// 		freqCount[c] = append(freqCount[c], r)
// 	}
// 	sort.Sort(sort.Reverse(sort.IntSlice(freqCountKeys)))
// 	result := make([]rune, 0, len(s))
// 	for _, count := range freqCountKeys {
// 		for _, r := range freqCount[count] {
// 			for i := 0; i < count; i++ {
// 				result = append(result, r)
// 			}
// 		}
// 	}
// 	return string(result)
// }
