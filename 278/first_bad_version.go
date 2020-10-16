// 278. First Bad Version
// https://leetcode.com/problems/first-bad-version/
package main

import (
	"sort"

	"github.com/k0kubun/pp"
)

// May LeetCoding Challenge Day 1
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3316/
func main() {
	pp.Println("=========================================")
	pp.Println("No local tests available.")
	pp.Println("=========================================")
}

/**
 * Forward declaration of isBadVersion API.
 * @param    version your guess about first bad version
 * @return   true if current version is bad
 *           false if current version is good
 * func isBadVersion(version int) bool;
 */

/*
 * 2. Using `sort.Search`
 */
func firstBadVersion(n int) int {
	return sort.Search(n+1, func(i int) bool { return isBadVersion(i) })
}

/*
 * 1. My binary search
 */
// func firstBadVersion(n int) int {
// 	leastGoodVersion, leastBadVersion := 0, n
// 	for leastGoodVersion != leastBadVersion-1 {
// 		if testTargetVersion := leastGoodVersion + (leastBadVersion-leastGoodVersion)>>1; isBadVersion(testTargetVersion) {
// 			leastBadVersion = testTargetVersion
// 		} else {
// 			leastGoodVersion = testTargetVersion
// 		}
// 	}
// 	return leastBadVersion
// }
