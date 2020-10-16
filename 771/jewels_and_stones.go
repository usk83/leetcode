// 771. Jewels and Stones
// https://leetcode.com/problems/jewels-and-stones/
package main

import (
	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 2
// https://leetcode.com/explore/featured/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3317/
func main() {
	pp.Println("=========================================")
	pp.Println(numJewelsInStones("aA", "aAAbbbb"))
	pp.Println(3)
	pp.Println("=========================================")
	pp.Println(numJewelsInStones("z", "ZZ"))
	pp.Println(0)
	pp.Println("=========================================")
}

func numJewelsInStones(jewels string, stones string) int {
	var jewelsList [58]bool // ASCII A(65)~z(122)
	for _, j := range jewels {
		jewelsList[j-'A'] = true
	}
	var count int
	for _, s := range stones {
		if jewelsList[s-'A'] {
			count++
		}
	}
	return count
}
