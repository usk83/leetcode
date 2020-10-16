// 383. Ransom Note
// https://leetcode.com/problems/ransom-note/
package main

import (
	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 3
// https://leetcode.com/explore/featured/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3318/
func main() {
	pp.Println("=========================================")
	pp.Println(canConstruct("a", "b"))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(canConstruct("aa", "ab"))
	pp.Println(false)
	pp.Println("=========================================")
	pp.Println(canConstruct("aa", "aab"))
	pp.Println(true)
	pp.Println("=========================================")
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	var letters [26]int
	for _, letter := range magazine {
		letters[letter-'a']++
	}
	for _, letter := range ransomNote {
		index := letter - 'a'
		if letters[index] == 0 {
			return false
		}
		letters[index]--
	}
	return true
}

/*
 * 1. first solution
 */
// func canConstruct(ransomNote string, magazine string) bool {
// 	target, source := [26]int{}, [26]int{}
// 	for i := range ransomNote {
// 		target[ransomNote[i]-'a']++
// 	}
// 	for i := range magazine {
// 		source[magazine[i]-'a']++
// 	}
// 	for i, count := range target {
// 		if count > source[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
