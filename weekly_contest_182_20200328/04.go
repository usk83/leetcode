// 5371. Find All Good Strings
// https://leetcode.com/contest/weekly-contest-182/problems/find-all-good-strings/
package main

import (
	"fmt"
	"math"

	"github.com/k0kubun/pp"
)

func main() {
	// pp.Println("=========================================")
	// pp.Println(findGoodStrings(2, "aa", "da", "b"))
	// pp.Println(51)
	// pp.Println("=========================================")
	// pp.Println(findGoodStrings(8, "leetcode", "leetgoes", "leet"))
	// pp.Println(0)
	// pp.Println("=========================================")
	// pp.Println(findGoodStrings(2, "gx", "gz", "x"))
	// pp.Println(2)
	pp.Println("=========================================")
	pp.Println(findGoodStrings(4, "aaaa", "abac", "x"))
	pp.Println(-1)
	pp.Println("=========================================")
	// pp.Println()
	// pp.Println()
	// pp.Println("=========================================")
}

const modint = 1000000000 + 7

// s1.length == n
// s2.length == n
// 1 <= n <= 500
// 1 <= evil.length <= 50
// All strings consist of lowercase English letters.
func findGoodStrings(n int, s1 string, s2 string, evil string) int {
	if s2 < s1 {
		return 0
	}

	afterCount := make([]int, n)
	afterCount[n-1] = int(s2[n-1] - s1[n-1])
	for i := n - 2; i >= 0; i-- {
		afterCount[i] = int(s2[i]-s1[i])*int(math.Pow(26, float64(n-i-1)))%modint + afterCount[i+1]
	}
	fmt.Println(afterCount)

	// afterCount[0] = int(s2[0] - s1[0])
	// for i := 1; i < n; i++ {
	// 	afterCount[i] = afterCount[i-1]%modint*26 + int(s2[i]-s1[i])
	// }
	// fmt.Println(afterCount)

	return -1
}
